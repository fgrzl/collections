package stack

func WithCapacity[T any](capacity int) func(*Stack[T]) {
	return func(s *Stack[T]) {
		s.items = make([]T, 0, capacity)
	}
}

// Stack is a generic LIFO stack.
type Stack[T any] struct {
	items []T
}

// NewStack initializes a stack with optional configuration.
func NewStack[T any](opts ...func(*Stack[T])) *Stack[T] {
	s := &Stack[T]{}
	for _, opt := range opts {
		opt(s)
	}
	if s.items == nil {
		s.items = make([]T, 0, 16)
	}
	return s
}

// Push adds an item to the top of the stack.
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the item from the top of the stack.
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if s.IsEmpty() {
		return zero, false
	}
	last := len(s.items) - 1
	item := s.items[last]
	s.items[last] = zero // Clear for GC
	s.items = s.items[:last]
	return item, true
}

// Peek returns the top item without removing it.
func (s *Stack[T]) Peek() (T, bool) {
	var zero T
	if s.IsEmpty() {
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

// IsEmpty checks if the stack has no items.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Length returns the number of items in the stack.
func (s *Stack[T]) Length() int {
	return len(s.items)
}

// Reset clears the stack while preserving capacity.
func (s *Stack[T]) Reset() {
	var zero T
	for i := range s.items {
		s.items[i] = zero
	}
	s.items = s.items[:0]
}
