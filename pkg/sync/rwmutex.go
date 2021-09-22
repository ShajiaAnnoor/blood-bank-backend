package sync

import "sync"

type ReadWriteMutex struct {
	mu sync.RWMutex
}

func (r *ReadWriteMutex) Read(f func()) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	f()
}

func (r *ReadWriteMutex) Write(f func()) {
	r.mu.Lock()
	defer r.mu.Unlock()

	f()
}
