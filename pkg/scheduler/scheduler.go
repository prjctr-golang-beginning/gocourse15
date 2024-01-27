package scheduler

import (
	"context"
	"github.com/samber/do"
)

type Task interface {
	TimeType() TimeType
	Expression() string
	Operation(ctx context.Context) func()
}

type TimeType uint8

const (
	Every TimeType = iota
	Cron
)

func (tt TimeType) String() string {
	switch tt {
	case Every:
		return `every`
	case Cron:
		return `cron`
	default:
		return `unknown`
	}
}

func ProvideScheduler(i *do.Injector) (*Scheduler, error) {
	return NewScheduler(i), nil
}

func NewScheduler(_ *do.Injector) *Scheduler {
	return &Scheduler{done: make(chan struct{})}
}

type Scheduler struct {
	done chan struct{}
}

func (r *Scheduler) Manage(_ context.Context, tasks ...Task) error {
	for i := range tasks {
		switch tasks[i].TimeType() {
		case Every:

		case Cron:

		}
	}

	<-r.done

	return nil
}

func (r *Scheduler) Shutdown() {
	close(r.done)
}
