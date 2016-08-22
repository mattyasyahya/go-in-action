package concurrency

import (
	"fmt"
	"runtime"
	"sync"
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

var counter int

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
