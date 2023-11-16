package server

import (
	"fmt"
	"sync"
	"testing"
)

func TestEnqueue(t *testing.T) {
	ms := NewMemoryStore()

	expectedSize := 1

	index, err := ms.Enqueue([]byte("Item"))
	if err != nil {
		t.Errorf("Enqueue returned error: %v", err)
	}

	if index != expectedSize-1 {
		t.Errorf("Enqueue returned the wrong index. Expected %d, got %d", expectedSize-1, index)
	}

	if len(ms.data) != expectedSize {
		t.Errorf("Size of data slice is incorrect after Enqueue. Expected %d, got %d", expectedSize, len(ms.data))
	}
}

func TestDequeue(t *testing.T) {
	ms := NewMemoryStore()

	dequeuedItem := ms.Dequeue()
	if dequeuedItem != nil {
		t.Error("Dequeue on an empty queue returned non-nil value, expected nil")
	}
}

func TestSize(t *testing.T) {
	ms := NewMemoryStore()

	// Test Size on an empty queue
	if size := ms.Size(); size != 0 {
		t.Errorf("Size on an empty queue is incorrect. Expected 0, got %d", size)
	}

	ms.Enqueue([]byte("Item"))
	if size := ms.Size(); size != 1 {
		t.Errorf("Size of queue is incorrect. Expected 1, got %d", size)
	}
}

func TestFront(t *testing.T) {
	ms := NewMemoryStore()

	frontItem := ms.Front()
	if frontItem != nil {
		t.Error("Front on an empty queue returned non-nil value, expected nil")
	}
}

func TestIsEmpty(t *testing.T) {
	ms := NewMemoryStore()

	if isEmpty := ms.IsEmpty(); !isEmpty {
		t.Error("IsEmpty on an empty queue returned false, expected true")
	}
}

func TestConcurrency(t *testing.T) {
	ms := NewMemoryStore()

	numGoroutines := 100

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(index int) {
			defer wg.Done()

			// Enqueue
			_, err := ms.Enqueue([]byte(fmt.Sprintf("Item%d", index)))
			if err != nil {
				t.Errorf("Enqueue returned an error: %v", err)
			}

			// Dequeue
			_ = ms.Dequeue()
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Check the final size of the queue
	expectedSize := 0
	if size := ms.Size(); size != expectedSize {
		t.Errorf("Final size of the queue is incorrect. Expected %d, got %d", expectedSize, size)
	}
}
