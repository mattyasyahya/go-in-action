package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

// CreateChannel simple channel
func CreateChannel() {
	channels := make(chan string, 10)

	channels <- "eko"
	channels <- "kurniawan"
	channels <- "khannedy"

	eko := <-channels
	fmt.Println(eko)

	kurniawan := <-channels
	fmt.Println(kurniawan)

	khannedy := <-channels
	fmt.Println(khannedy)
}

// TennisGame play tennis using channel
func TennisGame() {
	rand.Seed(time.Now().UnixNano())

	wg.Add(2)

	court := make(chan int)

	go player("Nadal", court)
	go player("Djokovic", court)

	court <- 1

	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court

		if !ok {
			fmt.Println("Player", name, "Won")
			return
		}

		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Println("Player", name, "Missed")
			close(court)
			return
		}

		fmt.Println("Player", name, "Hit", ball)
		ball++

		court <- ball
	}
}
