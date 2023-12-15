package task2

import (
	"fmt"
	"sync"
)

type LogBuffer struct {
	sync.Mutex
	buffer []string
}

func NewLogBuffer() *LogBuffer {
	return new(LogBuffer)
}

func (lb *LogBuffer) WriteLog(message string) {
	lb.Lock()
	defer lb.Unlock()

	lb.buffer = append(lb.buffer, message)
}

func (lb *LogBuffer) Print() {
	fmt.Println(lb.buffer)
}

func Start() {
	var wg sync.WaitGroup
	logBuffer := NewLogBuffer()
	messages := []string{
		"Какое-то сообщение",
		"Новое сообщение",
		"Еще сообщение",
		"test1",
		"test2",
		"test3",
		"test4",
		"test5",
	}

	for _, message := range messages {
		wg.Add(1)
		go func(msg string) {
			defer wg.Done()
			logBuffer.WriteLog(msg)
		}(message)
	}

	wg.Wait()
	logBuffer.Print() // ~[Какое-то сообщение Новое сообщение Еще сообщение test3 test2 test5 test4 test1]
}
