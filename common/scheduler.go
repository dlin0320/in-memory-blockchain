package common

import (
	"time"

	"github.com/reactivego/scheduler"
)

type Scheduler struct {
	scheduler.ConcurrentScheduler
}

type Task struct {
	Time float64
	Job  func()
}

func NewTask(time float64, job func()) *Task {
	return &Task{time, job}
}

func NewScheduler() *Scheduler {
	return &Scheduler{scheduler.Goroutine}
}

func (s *Scheduler) Schedule(t *Task) {
	s.ScheduleFuture(time.Duration(t.Time*float64(time.Second)), t.Job)
}
