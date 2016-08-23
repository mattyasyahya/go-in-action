package concurrency

import "fmt"

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
