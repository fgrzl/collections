package collections

type HashSet[T comparable] struct {
	data map[T]struct{}
}

func NewHashSet[T comparable]() HashSet[T] {
	return HashSet[T]{
		data: make(map[T]struct{}),
	}
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
