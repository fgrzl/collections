package concurrenthashset

import "sync"

// ConcurrentHashSet is a thread-safe version of HashSet.
type ConcurrentHashSet[T comparable] struct {
	mu   sync.RWMutex
	data map[T]struct{}
}

// ConcurrentHashSetOption defines a configuration function for ConcurrentHashSet.
type ConcurrentHashSetOption[T comparable] func(*ConcurrentHashSet[T])

// WithCapacity initializes the internal map with a given capacity.
func WithCapacity[T comparable](capacity int) ConcurrentHashSetOption[T] {
	return func(h *ConcurrentHashSet[T]) {
		h.data = make(map[T]struct{}, capacity)
	}
}

// NewConcurrentHashSet creates a new ConcurrentHashSet with optional configuration.
func NewConcurrentHashSet[T comparable](opts ...ConcurrentHashSetOption[T]) ConcurrentHashSet[T] {
	h := ConcurrentHashSet[T]{data: make(map[T]struct{})}
	for _, opt := range opts {
		opt(&h)
	}
	return h
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
