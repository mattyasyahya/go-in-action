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
	runtime.GOMAXPROCS(runtime.NumCPU())

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
