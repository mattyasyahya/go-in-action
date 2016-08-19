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
}
