package main

import (
	"mockery-test/handler"
	"mockery-test/logger"
)

func main() {
	logger := logger.NewLogger()
	printHandler := handler.NewPrintHandler(logger)

	printHandler.Run()
}
