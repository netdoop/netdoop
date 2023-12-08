package omc

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/netdoop/netdoop/store"
	"github.com/netdoop/netdoop/utils"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var defaultTaskServer *TaskServer
var defaultTaskServerOnce sync.Once

func GetTaskServer() *TaskServer {
	defaultTaskServerOnce.Do(func() {
		defaultTaskServer = NewTaskServer()
	})
	return defaultTaskServer
}

type TaskServer struct {
	sync.RWMutex
	logger *zap.Logger

	db       *gorm.DB
	sessions map[uint]*TaskSession

	shutdown     bool
	shutdownCh   chan struct{}
	shutdownLock sync.Mutex
}

func NewTaskServer() *TaskServer {
	return &TaskServer{
		db:         store.GetDB(),
		sessions:   make(map[uint]*TaskSession),
		shutdownCh: make(chan struct{}),
		logger:     utils.GetLogger().Named("tasksrv"),
	}
}

func (s *TaskServer) Close() {
	s.shutdownLock.Lock()
	defer s.shutdownLock.Unlock()

	if s.shutdown {
		return
	}
	s.logger.Debug("server close")
	s.shutdown = true
	close(s.shutdownCh)
}

func (s *TaskServer) Run() {
	s.logger.Debug("server running")
	defer s.logger.Debug("server stopped")

	var wg sync.WaitGroup
	defer wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()

		tasks, err := LoadAllTasks(s.db, "")
		if err != nil {
			s.logger.Fatal("start server", zap.Error(err))
		}
		for _, task := range tasks {
			s.StartTask(task)
		}
	}()

	<-s.shutdownCh

	for _, sess := range s.sessions {
		s.StopTask(sess.Task)
	}
}

func (s *TaskServer) ClearTask(task *Task) {
	s.Lock()
	defer s.Unlock()
	if sess, ok := s.sessions[task.ID]; ok {
		delete(s.sessions, task.ID)
		sess.Stop()
	}
}

func (s *TaskServer) StopTask(task *Task) {
	s.Lock()
	defer s.Unlock()
	if sess, ok := s.sessions[task.ID]; ok {
		sess.Stop()
	}
}

func (s *TaskServer) StartTask(task *Task) {
	s.Lock()
	defer s.Unlock()
	s.logger.Warn("task", zap.Uint("id", task.ID), zap.String("name", task.TaskName), zap.Any("task", *task))

	if sess, ok := s.sessions[task.ID]; ok {
		sess.Start()
		return
	}
	switch task.TaskType {
	case "NBI:HTTP:POST:DeviceAlarm":
		sess := CreateTaskSession(s.db, task, HandleHttpPost)
		s.sessions[task.ID] = sess
		sess.Start()
	default:
		return
	}
}

func (s *TaskServer) GetTaskSessionsBySource(source string) []*TaskSession {
	out := []*TaskSession{}
	for _, sess := range s.sessions {
		if parts := strings.Split(sess.Task.TaskType, ":"); len(parts) == 4 {
			if parts[3] == source {
				out = append(out, sess)
			}
		}
	}
	return out
}

func (s *TaskServer) PushDeviceAlarm(alarm *DeviceAlarm) {
	event := DeviceAlarmEvent{
		OUI:                   alarm.Oui,
		ProductClass:          alarm.ProductClass,
		SerialNumber:          alarm.SerialNumber,
		AlarmIdentifier:       alarm.AlarmIdentifier,
		EventType:             alarm.EventType,
		ManagedObjectInstance: alarm.ManagedObjectInstance,
		ProbableCause:         alarm.ProbableCause,
		SpecificProblem:       alarm.SpecificProblem,
		PerceivedSeverity:     alarm.PerceivedSeverity,
		AdditionalText:        alarm.AdditionalText,
		AdditionalInformation: alarm.AdditionalInformation,
	}
	if alarm.AlarmCleared {
		event.NotificationType = "ClearedAlarm"
		event.EventTime = alarm.AlarmClearedTime.Format()
	} else if alarm.AlarmChangedTime != alarm.AlarmRaisedTime {
		event.NotificationType = "ChangedAlarm"
		event.EventTime = alarm.AlarmChangedTime.Format()
	} else {
		event.NotificationType = "CreatedAlarm"
		event.EventTime = alarm.AlarmRaisedTime.Format()
	}
	s.logger.Warn("push job", zap.Any("event", event))
	sessions := s.GetTaskSessionsBySource("DeviceAlarm")
	for _, sess := range sessions {
		sess.PushJob(&event)
	}
}

type DeviceAlarmEvent struct {
	OUI                   string
	ProductClass          string
	SerialNumber          string
	EventTime             string
	AlarmIdentifier       string
	NotificationType      string
	ManagedObjectInstance string
	EventType             string
	ProbableCause         string
	SpecificProblem       string
	PerceivedSeverity     string
	AdditionalText        string
	AdditionalInformation string
}

func HandleHttpPost(task *Task, data []any) error {
	content, err := json.Marshal(data)
	if err != nil {
		return errors.Wrap(err, "marshal content for http post request")
	}
	retry := task.RetryTimes
	retryInterval := task.RetryInterval
	if retryInterval <= 0 {
		retryInterval = 1
	}
	url := task.MetaData.GetString("ServerUrl")
	for {
		resp, err := http.Post(url, "application/json", bytes.NewReader(content))
		utils.GetLogger().Debug("http post",
			zap.String("url", url),
			zap.Int("code", resp.StatusCode),
		)
		if err == nil && resp.StatusCode == 200 {
			break
		}
		if retry > 1 {
			retry -= 1
			time.Sleep(time.Second * time.Duration(retryInterval))
		}
	}
	return nil
}
