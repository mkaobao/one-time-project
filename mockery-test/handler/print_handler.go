package handler

import (
	"fmt"
	"mockery-test/logger"
)

type PrintHandler interface {
	Run()
}

var _ PrintHandler = (*PrintHandlerImpl)(nil)

func NewPrintHandler(logger logger.Logger) *PrintHandlerImpl {
	return &PrintHandlerImpl{
		logger: logger,
	}
}

type PrintHandlerImpl struct {
	logger logger.Logger
}

func (p *PrintHandlerImpl) Run() {
	for i := 0; i < 10; i++ {
		p.logger.Log(fmt.Sprintf("Hello, %v", i))
	}
}
