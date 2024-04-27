package Promocodes

import (
	"sync"
	"time"
)

type timerOnce struct {
	mu    sync.Mutex
	ch    chan struct{}
	timer *time.Timer
	once  sync.Once
	paid  bool
}

func newTimerOnce() *timerOnce {
	return &timerOnce{
		ch:   make(chan struct{}),
		paid: true,
	}
}

func (t *timerOnce) run(d time.Duration, f func()) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.timer = time.AfterFunc(d, func() {
		f()
		close(t.ch)
	})
}
