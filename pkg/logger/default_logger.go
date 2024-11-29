package logger

import "log"

type DefaultLogger struct{}

func (l *DefaultLogger) Info(msg string) {
	log.Println("INFO: ", msg)
}

func (l *DefaultLogger) Error(msg string) {
	log.Println("ERROR: ", msg)
}
