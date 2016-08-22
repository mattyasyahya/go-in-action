package concurrency

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// GetCPUCoreNum get cpu number
func GetCPUCoreNum() {
	num := runtime.NumCPU()
	fmt.Println(num)
}

// OneLogical run goroutine in one logical procs
func OneLogical() {
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Println("")
		}
	}()

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Println("")
		}
	}()

	fmt.Println("Waiting to finish")
	wg.Wait()

	fmt.Println("Terminating program")
}

var wg sync.WaitGroup

func printPrime(prefix string) {
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}

// TimeSlice demonstration time slice goroutine in single thread
func TimeSlice() {
	runtime.GOMAXPROCS(1)

	wg.Add(2)

	fmt.Println("Create Goroutines")
	go printPrime("A")
	go printPrime("B")
	fmt.Println("Waiting To Finish")

	wg.Wait()
	fmt.Println("Terminating Goroutines")
}

var counter int32

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		value := counter

		runtime.Gosched()

		value++

		counter = value
	}
}

// RaceCondition demonstration race condition
func RaceCondition() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Println("Final counter ", counter)
}

func safeIncCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		atomic.AddInt32(&counter, 1)
		runtime.Gosched()
	}
}

// SafeRaceCondition demonstration safe race condition with atomic
func SafeRaceCondition() {
	counter = 0

	wg.Add(2)

	go safeIncCounter(1)
	go safeIncCounter(2)

	wg.Wait()
	fmt.Println("Final counter ", counter)
}

var shutdown int64

func doWork(name string) {
	defer wg.Done()

	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)

		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}

// AtomicLoadAndStore demonstration atomic load and store operation
func AtomicLoadAndStore() {
	wg.Add(2)

	go doWork("A")
	go doWork("B")

	time.Sleep(1 * time.Second)

	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)

	wg.Wait()
}

var mutex sync.Mutex

func mutexIncCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {

		mutex.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock()

	}
}

// MutexRaceCondition demonstration safe race condition using mutex
func MutexRaceCondition() {
	counter = 0
	wg.Add(2)

	go mutexIncCounter(1)
	go mutexIncCounter(2)

	wg.Wait()
	fmt.Println("Final counter ", counter)
}
