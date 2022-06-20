package common

import (
	"time"

	"github.com/reactivego/scheduler"
)

type Scheduler interface {
	Schedule(t *Task)
	Wait()
}

type SingleScheduler struct {
	scheduler.Scheduler
}

type RepeatedScheduler struct {
	scheduler.Scheduler
}

type Task struct {
	Time float64
	Job  func()
}

func (s RepeatedScheduler) Schedule(t *Task) {
	s.ScheduleRecursive(func(again func()) {
		time.Sleep(time.Duration(t.Time * float64(time.Second)))
		t.Job()
		again()
	})
}

func (s SingleScheduler) Schedule(t *Task) {
	s.ScheduleFuture(time.Duration(t.Time*float64(time.Second)), t.Job)
}

func NewTask(time float64, job func()) *Task {
	return &Task{time, job}
}

func NewScheduler(repeated bool, concurrent bool) Scheduler {
	var s scheduler.Scheduler
	if concurrent {
		s = scheduler.Goroutine
	} else {
		s = scheduler.New()
	}
	if repeated {
		return RepeatedScheduler{s}
	}
	return SingleScheduler{s}
}
