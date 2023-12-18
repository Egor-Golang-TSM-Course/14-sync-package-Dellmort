package task4

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type Handler struct {
	id      int
	timeout time.Duration
	sync.Mutex
	status bool
}

func NewHandler(id int, timeout time.Duration) *Handler {
	return &Handler{
		id:      id,
		timeout: timeout,
	}
}

func (h *Handler) Listen(request <-chan string) {
	h.Lock()
	defer h.Unlock()

	h.status = true
	h.handle(request)
}

func (h *Handler) handle(request <-chan string) {
	req, ok := <-request
	if ok && req != "" {
		// Какая-то обработка
		fmt.Printf("обработчик №%d, %s\n", h.id, strings.ToUpper(req))
		<-time.After(h.timeout * time.Second)
	}
	h.status = false
}

func (h *Handler) Status() bool {
	h.Lock()
	defer h.Unlock()

	return h.status
}
