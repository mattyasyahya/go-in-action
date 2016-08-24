package loglib

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	// Trace logger
	Trace *log.Logger
	// Info logger
	Info *log.Logger
	// Warning logger
	Warning *log.Logger
	// Error logger
	Error *log.Logger
)

func initLogger() {
	file, err := os.OpenFile("logs/application.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)
	}

	Trace = log.New(ioutil.Discard, "TRACE:", log.Ldate|log.Lmicroseconds|log.Llongfile)
	Info = log.New(os.Stdout, "INFO:", log.Ldate|log.Lmicroseconds|log.Llongfile)
	Warning = log.New(os.Stdout, "WARNING:", log.Ldate|log.Lmicroseconds|log.Llongfile)
	Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR:", log.Ldate|log.Lmicroseconds|log.Llongfile)
}

// CustomeLogger sample custome logger
func CustomeLogger() {
	initLogger()

	Trace.Println("I have something standard to say")
	Info.Println("Special Information")
	Warning.Println("There is something you need to know about")
	Error.Println("Something has failed")
}
