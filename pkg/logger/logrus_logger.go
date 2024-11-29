package logger

import (
	log "github.com/sirupsen/logrus"
)

type LogrusLogger struct{}

func (l *LogrusLogger) Info(msg string) {
	log.Info(msg)
}

func (l *LogrusLogger) Error(msg string) {
	log.Error(msg)
}
