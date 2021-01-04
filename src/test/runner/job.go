package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

type JobRunner struct {
	 interrupt chan os.Signal
	 compete chan error
	 timeout <- chan time.Time
	 tasks []func(int)
}

var ErrTimeout = errors.New("received timeout")
var ErrInterrut = errors.New("received interrupt")

func (r *JobRunner) Add(tasks ...func(int))  {
	r.tasks = append(r.tasks, tasks...)
}

func (r *JobRunner) Start() error  {
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.compete <- r.run()
	}()

	select {
		case err := <- r.compete :
			return err
		case <- r.timeout:
			return ErrInterrut
	}
}

func (r *JobRunner) run() error {
	for id,task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrut
		}
		task(id)
	}
	return nil
}

func (r * JobRunner) gotInterrupt() bool {
	select {
		case <- r.interrupt:
			signal.Stop(r.interrupt)
			return true
	default:
		return false
	}
}