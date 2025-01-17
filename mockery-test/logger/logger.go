package logger

import "fmt"

type Logger interface {
	Log(message string) string
}

func NewLogger() *LoggerImpl {
	return &LoggerImpl{}
}

type LoggerImpl struct{}

func (l *LoggerImpl) Log(message string) string {
	output := "Log: " + message

	fmt.Println(output)

	return output
}
