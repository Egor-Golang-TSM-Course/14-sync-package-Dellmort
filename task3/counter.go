package task3

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type WebCounter struct {
	visits sync.Map
}

func New() *WebCounter {
	return new(WebCounter)
}

func (w *WebCounter) Increment(url string) {
	key, _ := w.visits.LoadOrStore(url, int32(0))
	count := key.(int32)
	w.visits.Swap(url, atomic.AddInt32(&count, 1))
}

func (w *WebCounter) GetVisits(url string) int {
	key, _ := w.visits.Load(url)
	if value, ok := key.(int32); ok {
		return int(atomic.LoadInt32(&value))
	}
	return 0
}

func (w *WebCounter) Print() {
	w.visits.Range(func(key, value any) bool {
		val := w.GetVisits(key.(string))
		fmt.Println(val)
		return true
	})
}

func Start() {
	var wg sync.WaitGroup
	wb := New()
	wg.Add(2)
	go func() {
		defer wg.Done()
		wb.Increment("google.com")
		wb.Increment("google.com")
		wb.Increment("google.com")
		wb.Increment("google.com")
	}()
	go func() {
		defer wg.Done()
		wb.Increment("yandex.ru")
		wb.Increment("yandex.ru")
		wb.Increment("yandex.ru")
	}()
	wg.Wait()

	wb.Print()
}
