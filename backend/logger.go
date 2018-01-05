package main

import (
	"log"
	"os"
)

type Logger struct {
	info *log.Logger
	warn *log.Logger
	err  *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime),
		log.New(os.Stdout, "WARNING ", log.Ldate|log.Ltime),
		log.New(os.Stderr, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
