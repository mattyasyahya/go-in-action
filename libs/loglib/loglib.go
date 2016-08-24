package loglib

import "log"

// Main sample log
func Main() {
	log.Println("Hello world")
	log.Fatalln("Ups, error")
	log.Panicln("Ups, panic")
}
