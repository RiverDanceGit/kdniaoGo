package kdniaoGo

import (
	"log"
)

type KdniaoLoggerInterface interface {
	Error(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})
	Warn(args ...interface{})
	Panic(args ...interface{})
}

type KdniaoLogger struct {
}

func NewKdniaoLogger() KdniaoLoggerInterface {
	return &KdniaoLogger{}
}

func (logs *KdniaoLogger) Error(args ...interface{}) {
	log.Fatal(args)
}

func (logs *KdniaoLogger) Info(args ...interface{}) {
	log.Println("INFO", args)
}

func (logs *KdniaoLogger) Debug(args ...interface{}) {
	log.Println("DEBUG", args)
}

func (logs *KdniaoLogger) Warn(args ...interface{}) {
	log.Println("WARN", args)
}

func (logs *KdniaoLogger) Panic(args ...interface{}) {
	log.Panic(args)
}
