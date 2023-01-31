package logger

import "go.uber.org/zap"

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger() *Logger {
	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	return &Logger{
		log.Sugar(),
	}
}
