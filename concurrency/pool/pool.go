package pool

import (
	"errors"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// Pool type
type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

// Acquirer new pool
func (pool *Pool) Acquirer() (io.Closer, error) {
	select {
	case r, ok := <-pool.resources:
		log.Println("Acquire: Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Acquire: New Resource")
		return pool.factory()
	}
}

// Release resource pool
func (pool *Pool) Release(resource io.Closer) {
	pool.m.Lock()
	defer pool.m.Unlock()

	if pool.closed {
		resource.Close()
		return
	}

	select {
	case pool.resources <- resource:
		log.Println("Release: In Queue")

	default:
		log.Println("Release: Closing")
		resource.Close()
	}
}

// Close the pool
func (pool *Pool) Close() {
	pool.m.Lock()
	defer pool.m.Unlock()

	if pool.closed {
		return
	}

	pool.closed = true
	close(pool.resources)

	for resource := range pool.resources {
		resource.Close()
	}
}

// ErrPoolClosed returned when an closed pool
var ErrPoolClosed = errors.New("Pool has been closed")

// New create new pool
func New(fn func() (io.Closer, error), size int) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Size value too small")
	}

	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

const (
	maxGoRoutines   = 25
	pooledResources = 2
)

// DbConnection sample db connection
type DbConnection struct {
	ID int32
}

// Close db connection
func (db *DbConnection) Close() error {
	log.Println("Close: connection", db.ID)
	return nil
}

var idCounter int32

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)
	return &DbConnection{id}, nil
}

// Main program for pool
func Main() {
	var wg sync.WaitGroup
	wg.Add(maxGoRoutines)

	p, err := New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	for query := 0; query < maxGoRoutines; query++ {
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()
	log.Println("Shutdown program")
	p.Close()

}

func performQueries(query int, p *Pool) {
	conn, err := p.Acquirer()
	if err != nil {
		log.Println(err)
		return
	}

	defer p.Release(conn)

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*DbConnection).ID)
}
