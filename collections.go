package collections

// Re-exports from stack package
import (
	"github.com/fgrzl/collections/concurrenthashset"
	"github.com/fgrzl/collections/hashset"
	"github.com/fgrzl/collections/queue"
	"github.com/fgrzl/collections/stack"
)

// Stack re-exports
type Stack[T any] = stack.Stack[T]

// NewStack creates a new stack with optional configuration.
func NewStack[T any](opts ...func(*Stack[T])) *Stack[T] {
	return stack.NewStack(opts...)
}

// NewStackWithCapacity creates a new stack with preallocated capacity.
func NewStackWithCapacity[T any](capacity int) *Stack[T] {
	return stack.NewStack(stack.WithCapacity[T](capacity))
}

// StackWithCapacity returns an option that preallocates the stack with the given capacity.
func StackWithCapacity[T any](capacity int) func(*Stack[T]) {
	return stack.WithCapacity[T](capacity)
}

// Queue re-exports
type Queue[T any] = queue.Queue[T]

// NewQueue creates a new queue with optional configuration.
func NewQueue[T any](opts ...func(*Queue[T])) *Queue[T] {
	return queue.NewQueue(opts...)
}

// NewQueueWithCapacity creates a new queue with preallocated capacity.
func NewQueueWithCapacity[T any](capacity int) *Queue[T] {
	return queue.NewQueue(queue.WithCapacity[T](capacity))
}

// QueueWithCapacity returns an option that preallocates the queue with the given capacity.
func QueueWithCapacity[T any](capacity int) func(*Queue[T]) {
	return queue.WithCapacity[T](capacity)
}

// HashSet re-exports
type HashSet[T comparable] = hashset.HashSet[T]
type HashSetOption[T comparable] = hashset.HashSetOption[T]

// NewHashSet creates a new HashSet with optional configuration.
func NewHashSet[T comparable](opts ...HashSetOption[T]) *HashSet[T] {
	return hashset.NewHashSet(opts...)
}

// NewHashSetWithCapacity creates a new HashSet with preallocated capacity.
func NewHashSetWithCapacity[T comparable](capacity int) *HashSet[T] {
	return hashset.NewHashSet(hashset.WithCapacity[T](capacity))
}

// HashSetWithCapacity initializes the internal map with a given capacity.
func HashSetWithCapacity[T comparable](capacity int) HashSetOption[T] {
	return hashset.WithCapacity[T](capacity)
}

// ConcurrentHashSet re-exports
type ConcurrentHashSet[T comparable] = concurrenthashset.ConcurrentHashSet[T]
type ConcurrentHashSetOption[T comparable] = concurrenthashset.ConcurrentHashSetOption[T]

// NewConcurrentHashSet creates a new ConcurrentHashSet with optional configuration.
func NewConcurrentHashSet[T comparable](opts ...ConcurrentHashSetOption[T]) *ConcurrentHashSet[T] {
	return concurrenthashset.NewConcurrentHashSet(opts...)
}

// NewConcurrentHashSetWithCapacity creates a new ConcurrentHashSet with preallocated capacity.
func NewConcurrentHashSetWithCapacity[T comparable](capacity int) *ConcurrentHashSet[T] {
	return concurrenthashset.NewConcurrentHashSet(concurrenthashset.WithCapacity[T](capacity))
}

// ConcurrentHashSetWithCapacity initializes the internal map with a given capacity.
func ConcurrentHashSetWithCapacity[T comparable](capacity int) ConcurrentHashSetOption[T] {
	return concurrenthashset.WithCapacity[T](capacity)
}

// Type inference helpers for common usage patterns

// StackOf creates a new stack and initializes it with the provided values.
// The type is inferred from the values provided.
func StackOf[T any](values ...T) *Stack[T] {
	s := NewStack[T]()
	for _, v := range values {
		s.Push(v)
	}
	return s
}

// QueueOf creates a new queue and initializes it with the provided values.
// The type is inferred from the values provided.
func QueueOf[T any](values ...T) *Queue[T] {
	q := NewQueue[T]()
	for _, v := range values {
		q.Enqueue(v)
	}
	return q
}

// HashSetOf creates a new HashSet and initializes it with the provided values.
// The type is inferred from the values provided.
func HashSetOf[T comparable](values ...T) *HashSet[T] {
	hs := NewHashSet[T]()
	for _, v := range values {
		hs.Add(v)
	}
	return hs
}

// ConcurrentHashSetOf creates a new ConcurrentHashSet and initializes it with the provided values.
// The type is inferred from the values provided.
func ConcurrentHashSetOf[T comparable](values ...T) *ConcurrentHashSet[T] {
	chs := NewConcurrentHashSet[T]()
	for _, v := range values {
		chs.Add(v)
	}
	return chs
}

// Fluent builders for more complex configurations

// StackBuilder provides a fluent interface for building stacks with multiple options.
type StackBuilder[T any] struct {
	capacity *int
	values   []T
}

// NewStackBuilder creates a new StackBuilder.
func NewStackBuilder[T any]() *StackBuilder[T] {
	return &StackBuilder[T]{}
}

// WithCapacity sets the initial capacity for the stack.
func (sb *StackBuilder[T]) WithCapacity(capacity int) *StackBuilder[T] {
	sb.capacity = &capacity
	return sb
}

// WithValues initializes the stack with the provided values.
func (sb *StackBuilder[T]) WithValues(values ...T) *StackBuilder[T] {
	sb.values = append(sb.values, values...)
	return sb
}

// Build creates the stack with all configured options.
func (sb *StackBuilder[T]) Build() *Stack[T] {
	var stackOpts []func(*Stack[T])
	if sb.capacity != nil {
		stackOpts = append(stackOpts, StackWithCapacity[T](*sb.capacity))
	}
	s := NewStack(stackOpts...)
	for _, v := range sb.values {
		s.Push(v)
	}
	return s
}

// QueueBuilder provides a fluent interface for building queues with multiple options.
type QueueBuilder[T any] struct {
	capacity *int
	values   []T
}

// NewQueueBuilder creates a new QueueBuilder.
func NewQueueBuilder[T any]() *QueueBuilder[T] {
	return &QueueBuilder[T]{}
}

// WithCapacity sets the initial capacity for the queue.
func (qb *QueueBuilder[T]) WithCapacity(capacity int) *QueueBuilder[T] {
	qb.capacity = &capacity
	return qb
}

// WithValues initializes the queue with the provided values.
func (qb *QueueBuilder[T]) WithValues(values ...T) *QueueBuilder[T] {
	qb.values = append(qb.values, values...)
	return qb
}

// Build creates the queue with all configured options.
func (qb *QueueBuilder[T]) Build() *Queue[T] {
	var queueOpts []func(*Queue[T])
	if qb.capacity != nil {
		queueOpts = append(queueOpts, QueueWithCapacity[T](*qb.capacity))
	}
	q := NewQueue(queueOpts...)
	for _, v := range qb.values {
		q.Enqueue(v)
	}
	return q
}
