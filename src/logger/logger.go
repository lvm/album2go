package logger

import (
	"log"
)

var verbose bool

func SetVerbose(v bool) {
	verbose = v
}

func Info(msg string) {
	if verbose {
		log.Printf(msg)
	}
}

func Error(format string, err error) {
	log.Fatalf("Error: %v", err)
}
