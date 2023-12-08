package utils

import (
	"sync"
	"time"

	"github.com/RussellLuo/timingwheel"
	"go.uber.org/zap"
)

type TimingServer interface {
	Run()
	Stop()
	StartTimer(key string, d time.Duration, f func())
	StopTimer(key string)
}

type BaseTimingServer struct {
	logger    *zap.Logger
	tw        *timingwheel.TimingWheel
	timers    map[string]*timingwheel.Timer
	timersMtx sync.RWMutex
	quit      chan bool
}

func NewBaseTimingServer(tick time.Duration, size int64) (TimingServer, error) {
	s := &BaseTimingServer{
		quit:   make(chan bool),
		timers: map[string]*timingwheel.Timer{},
	}
	s.tw = timingwheel.NewTimingWheel(tick, size)
	s.logger = GetLogger().Named("timingsrv")
	return s, nil
}

func (s *BaseTimingServer) Run() {
	s.logger.Debug("server running")
	defer s.logger.Debug("server stopped")
	s.tw.Start()
	<-s.quit
	s.tw.Stop()
}

func (s *BaseTimingServer) Stop() {
	s.logger.Debug("server stop")
	s.quit <- true
}

func (s *BaseTimingServer) StartTimer(key string, d time.Duration, f func()) {
	s.stopTimer(key)
	if f == nil {
		return
	}
	t := s.tw.AfterFunc(d, func() {
		f()
	})
	s.timersMtx.Lock()
	s.timers[key] = t
	s.timersMtx.Unlock()
}

func (s *BaseTimingServer) StopTimer(key string) {
	s.stopTimer(key)
}

func (s *BaseTimingServer) stopTimer(key string) {
	s.timersMtx.Lock()
	defer s.timersMtx.Unlock()
	t, ok := s.timers[key]
	if ok && t != nil {
		t.Stop()
		delete(s.timers, key)
	}
}

var (
	tsrv TimingServer
	once sync.Once
)

func GetTimingServer() TimingServer {
	once.Do(func() {
		var err error
		tsrv, err = NewBaseTimingServer(time.Millisecond, 20)
		if err != nil {
			panic(err)
		}
		go tsrv.Run()
	})
	return tsrv
}
