package work

import (
	"log"
	"sync"
	"time"
)

// Worker interface
type Worker interface {
	Task()
}

// Pool struct
type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

// New create new pool with maxGoroutines
func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}

	p.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}

	return &p
}

// Run a worker
func (p *Pool) Run(w Worker) {
	p.work <- w
}

// Shutdown shutdown
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}

var names = []string{
	"steve",
	"bob",
	"mary",
	"theresa",
	"jason",
}

// NamePrinter struct
type NamePrinter struct {
	name string
}

// Task implementation for NamePrinter
func (n *NamePrinter) Task() {
	log.Println(n.name)
	time.Sleep(time.Millisecond * 10)
}

// Main program for work pattern
func Main() {
	p := New(2)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		for _, name := range names {
			np := NamePrinter{
				name: name,
			}

			go func() {
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()
	p.Shutdown()
}
