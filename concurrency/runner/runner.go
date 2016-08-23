package runner

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"
)

// ErrTimeout error for timeout
var ErrTimeout = errors.New("received timeout")

// ErrInterrupt error for interrupt
var ErrInterrupt = errors.New("received interrupt")

// Runner type
type Runner struct {
	interrupt chan os.Signal
	complete  chan error
	timeout   <-chan time.Time
	tasks     []func(int)
}

// Add task to runner
func (runner *Runner) Add(tasks ...func(int)) {
	runner.tasks = append(runner.tasks, tasks...)
}

// IsInterrupted check is interrupt
func (runner *Runner) IsInterrupted() bool {
	select {
	case <-runner.interrupt:
		signal.Stop(runner.interrupt)
		return true

	default:
		return false
	}
}

// Run tasks
func (runner *Runner) Run() error {
	for id, task := range runner.tasks {
		if runner.IsInterrupted() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

// Start runner
func (runner *Runner) Start() error {
	signal.Notify(runner.interrupt, os.Interrupt)

	go func() {
		runner.complete <- runner.Run()
	}()

	select {
	case err := <-runner.complete:
		return err
	case <-runner.timeout:
		return ErrTimeout
	}
}

// New create new Runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

const timeout = 3 * time.Second

// Main runner program
func Main() {
	log.Println("Starting work")

	runner := New(timeout)
	runner.Add(createTask(), createTask(), createTask())

	if err := runner.Start(); err != nil {
		switch err {
		case ErrTimeout:
			log.Println("Terminating due to timeout")
			os.Exit(1)
		case ErrInterrupt:
			log.Println("Terminating due to interrupt")
			os.Exit(1)
		}
	}
}

func createTask() func(int) {
	return func(id int) {
		log.Println("Processor task ", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
