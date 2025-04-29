package collections

import "sync"

// ConcurrentHashSet is a thread-safe version of HashSet.
type ConcurrentHashSet[T comparable] struct {
	mu   sync.RWMutex
	data map[T]struct{}
}

// NewConcurrentHashSet creates a new, empty ConcurrentHashSet.
func NewConcurrentHashSet[T comparable]() ConcurrentHashSet[T] {
	return ConcurrentHashSet[T]{
		data: make(map[T]struct{}),
	}
}

func (h *ConcurrentHashSet[T]) Add(item T) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.data[item] = struct{}{}
}

func (h *ConcurrentHashSet[T]) Remove(item T) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.data, item)
}

func (h *ConcurrentHashSet[T]) Contains(item T) bool {
	h.mu.RLock()
	defer h.mu.RUnlock()
	_, exists := h.data[item]
	return exists
}

func (h *ConcurrentHashSet[T]) IsEmpty() bool {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(h.data) == 0
}

func (h *ConcurrentHashSet[T]) Size() int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(h.data)
}

func (h *ConcurrentHashSet[T]) Clear() {
	h.mu.Lock()
	defer h.mu.Unlock()
	clear(h.data)
}

func (h *ConcurrentHashSet[T]) ToSlice() []T {
	h.mu.RLock()
	defer h.mu.RUnlock()

	slice := make([]T, 0, len(h.data))
	for item := range h.data {
		slice = append(slice, item)
	}
	return slice
}

func (h *ConcurrentHashSet[T]) ForEach(action func(T)) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	for item := range h.data {
		action(item)
	}
}
