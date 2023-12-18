package task4

import (
	"fmt"
	"time"
)

type Processor struct {
	request  chan string
	timeout  time.Duration
	handlers []*Handler
}

func NewProcessor(timeout time.Duration, request chan string, handler ...*Handler) *Processor {
	return &Processor{
		timeout:  timeout,
		request:  request,
		handlers: handler,
	}
}

func (h *Processor) Handle() {
	timer := time.After(h.timeout * time.Second)
	for {
		select {
		case <-timer:
			fmt.Println("Выхожу по таймауту")
			return
		default:
			for _, handler := range h.handlers {
				if handler.Status() {
					continue
				}
				go handler.Listen(h.request)
			}
		}
	}
}

func Start() {
	request := make(chan string)
	go func() {
		for i := 0; i < 15; i++ {
			request <- "send"
		}
		close(request)
	}()

	handler1 := NewHandler(1, 3)
	handler2 := NewHandler(2, 2)
	handler3 := NewHandler(3, 5)
	handler4 := NewHandler(4, 8)

	proc := NewProcessor(15, request, handler1, handler2, handler3, handler4)
	proc.Handle()
}
