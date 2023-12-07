package omc

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/netdoop/netdoop/utils"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"gorm.io/gorm"
)

type TaskHandleFunc func(task *Task, data []any) error

type TaskJob struct {
	Data any
}

type TaskSession struct {
	sync.Mutex
	Key    string
	db     *gorm.DB
	Task   *Task
	Status TaskStatus

	queuedJobs chan *TaskJob
	handleFunc TaskHandleFunc

	Process     int64
	LastTime    int64
	NextTime    int64
	CurrentRate int64
	Timeout     int64

	logger *zap.Logger
}

func CreateTaskSession(db *gorm.DB, task *Task, f TaskHandleFunc) *TaskSession {
	sess := TaskSession{
		Key:        fmt.Sprintf("%v-%v", task.ID, task.TaskType),
		db:         db,
		Task:       task,
		Status:     task.TaskStatus,
		handleFunc: f,
	}
	sess.queuedJobs = make(chan *TaskJob, task.ExecRate*10)
	sess.logger = utils.GetLogger().Named(sess.Key)
	return &sess
}

func (s *TaskSession) Stop() {
	s.Lock()
	defer s.Unlock()
	s.logger.Debug("task stop")
	s.stopNextExecution()
}

func (s *TaskSession) Start() {
	s.Lock()
	defer s.Unlock()
	s.logger.Debug("task start")
	s.startNextExecutionn()
}

func (s *TaskSession) PushJob(data any) {
	job := TaskJob{
		Data: data,
	}
	s.queuedJobs <- &job
}

func (s *TaskSession) setStatus(status TaskStatus) {
	if s.Status != status {
		s.Status = status
		if s.Status == TaskStatusFinish || s.Status == TaskStatusCancel {
			UpdateTaskStatus(s.db, "", s.Task, s.Status)
		}
	}
}

func (s *TaskSession) stopNextExecution() {
	tsrv := utils.GetTimingServer()
	tsrv.StopTimer(s.Key)
	s.setStatus(TaskStatusCancel)
}

func (s *TaskSession) startNextExecutionn() {
	if s.Status == TaskStatusFinish {
		return
	}
	if s.Status == TaskStatusCancel {
		return
	}

	d := s.getNextDuration()
	// s.logger.Debug("next duration", zap.Int("status", int(s.Status)), zap.Duration("duration", d))

	if d < 0 {
		s.setStatus(TaskStatusFinish)
		return
	}

	s.setStatus(TaskStatusWaiting)
	utils.GetTimingServer().StartTimer(s.Key, d, s.execute)
}

func (s *TaskSession) getNextDuration() time.Duration {
	ts := time.Now().Unix()
	nextTime := int64(0)
	interval := int64(s.Task.ExecInterval)
	if interval <= 0 {
		interval = 1
	}

	if s.Task.ExecMode == TaskExecModeSchedule {
		startTime := time.Time(s.Task.ExecStartTime).Unix()
		endTime := time.Time(s.Task.ExecEndTime).Unix()
		if startTime <= 0 {
			nextTime = ts + interval
		} else if startTime >= ts {
			nextTime = startTime
		} else {
			nextTime = ts + (interval - (ts-startTime)%interval)
		}
		if endTime > 0 && nextTime >= endTime {
			s.NextTime = -1
			return -1
		}
	} else {
		nextTime = ts + interval
	}

	if nextTime < ts {
		nextTime = ts
	}
	s.NextTime = nextTime
	return time.Duration(nextTime-ts) * time.Second
}

func (s *TaskSession) execute() {
	s.Lock()
	s.setStatus(TaskStatusRunning)
	s.Unlock()

	switch s.Task.TaskType {
	case "NBI:HTTP:POST:DeviceAlarm":
		if err := s.handleHttpPostDeviceAlarmJob(); err != nil {
			s.logger.Error("handle http post for device alarm", zap.Error(err))
		}
	default:
		return
	}

	s.Lock()
	defer s.Unlock()
	s.startNextExecutionn()
}

func (s *TaskSession) handleHttpPostDeviceAlarmJob() error {
	limit := int(s.Task.ExecRate)
	if limit <= 0 {
		limit = 1000
	}
	burst := limit / 10
	if burst <= 0 {
		burst = 1
	}

	r := rate.Every(time.Second / time.Duration(limit))
	limiter := rate.NewLimiter(r, burst)

	more := false
	size := len(s.queuedJobs)
	if size > 0 {
		s.logger.Info("queued", zap.Int("size", size))
		if size > burst {
			if size-burst > 0 {
				more = true
			}
			size = burst
		}
		data := make([]any, size)
		for i := 0; i < size; i++ {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			limiter.Wait(ctx)
			cancel()
			if s.Status == TaskStatusCancel {
				break
			}
			v := <-s.queuedJobs
			data[i] = v.Data
		}
		if err := s.handleFunc(s.Task, data); err != nil {
			return errors.Wrap(err, "hamdle task job")
		}
	}
	if more {
		time.Sleep(time.Millisecond * 50)
		return s.handleHttpPostDeviceAlarmJob()
	}
	return nil
}
