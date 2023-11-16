package server

import "sync"

type Queue interface {
	Enqueue([]byte) (int, error)
	Dequeue() []byte
	Front() []byte
	IsEmpty() bool
	Size() int
}

type MemoryStore struct {
	mu   sync.RWMutex
	data [][]byte
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make([][]byte, 0),
	}
}

func (ms *MemoryStore) Enqueue(b []byte) (int, error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	ms.data = append(ms.data, b)
	return len(ms.data) - 1, nil
}

func (ms *MemoryStore) Dequeue() []byte {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	if len(ms.data) == 0 {
		return nil
	}

	item := ms.data[0]
	ms.data = ms.data[1:]
	return item
}

func (ms *MemoryStore) Front() []byte {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	if len(ms.data) == 0 {
		return nil
	}
	return ms.data[0]
}

func (ms *MemoryStore) IsEmpty() bool {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	return len(ms.data) == 0
}

func (ms *MemoryStore) Size() int {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	return len(ms.data)
}
