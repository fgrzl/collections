package hashset

// HashSet is a simple set backed by a map.
type HashSet[T comparable] struct {
	data map[T]struct{}
}

// HashSetOption defines a configuration function for HashSet.
type HashSetOption[T comparable] func(*HashSet[T])

// WithCapacity initializes the internal map with a given capacity.
func WithCapacity[T comparable](capacity int) HashSetOption[T] {
	return func(h *HashSet[T]) {
		h.data = make(map[T]struct{}, capacity)
	}
}

// NewHashSet creates a new HashSet with optional configuration.
func NewHashSet[T comparable](opts ...HashSetOption[T]) HashSet[T] {
	h := HashSet[T]{data: make(map[T]struct{})}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

func (h HashSet[T]) Add(item T) {
	h.data[item] = struct{}{}
}

func (h HashSet[T]) Remove(item T) {
	delete(h.data, item)
}

func (h HashSet[T]) Contains(item T) bool {
	_, exists := h.data[item]
	return exists
}

func (h HashSet[T]) IsEmpty() bool {
	return len(h.data) == 0
}

func (h HashSet[T]) Size() int {
	return len(h.data)
}

func (h HashSet[T]) Clear() {
	clear(h.data)
}

func (h HashSet[T]) ToSlice() []T {
	slice := make([]T, 0, len(h.data))
	for item := range h.data {
		slice = append(slice, item)
	}
	return slice
}

func (h HashSet[T]) ForEach(action func(T)) {
	for item := range h.data {
		action(item)
	}
}
