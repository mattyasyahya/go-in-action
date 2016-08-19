package main

import (
	"log"

	"github.com/TechResearchID/go-in-action/math"
)

func init() {
	log.Println("Init program")
}

func main() {
	log.Println("Start program")
	log.Println(math.Max(1, 2))
	log.Println(math.Average(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
