package collections_test

import (
	"testing"

	"github.com/fgrzl/collections"
)

func TestReExports(t *testing.T) {
	// Test Stack re-export - original verbose way
	stack := collections.NewStack[int](collections.StackWithCapacity[int](10))
	stack.Push(1)
	stack.Push(2)

	if val, ok := stack.Pop(); !ok || val != 2 {
		t.Errorf("Stack re-export failed: expected 2, got %v", val)
	}

	// Test Queue re-export - original verbose way
	queue := collections.NewQueue[string](collections.QueueWithCapacity[string](5))
	queue.Enqueue("first")
	queue.Enqueue("second")

	if val, ok := queue.Dequeue(); !ok || val != "first" {
		t.Errorf("Queue re-export failed: expected 'first', got %v", val)
	}

	// Test HashSet re-export - original verbose way
	hashSet := collections.NewHashSet[int](collections.HashSetWithCapacity[int](20))
	hashSet.Add(1)
	hashSet.Add(2)

	if !hashSet.Contains(1) || !hashSet.Contains(2) {
		t.Error("HashSet re-export failed: values not found")
	}

	// Test ConcurrentHashSet re-export - original verbose way
	concurrentSet := collections.NewConcurrentHashSet[string](collections.ConcurrentHashSetWithCapacity[string](15))
	concurrentSet.Add("test")

	if !concurrentSet.Contains("test") {
		t.Error("ConcurrentHashSet re-export failed: value not found")
	}
}

func TestConvenienceConstructors(t *testing.T) {
	// Test convenience constructors - much cleaner!
	stack := collections.NewStackWithCapacity[int](10)
	stack.Push(42)

	if val, ok := stack.Pop(); !ok || val != 42 {
		t.Errorf("Stack convenience constructor failed: expected 42, got %v", val)
	}

	queue := collections.NewQueueWithCapacity[string](5)
	queue.Enqueue("convenience")

	if val, ok := queue.Dequeue(); !ok || val != "convenience" {
		t.Errorf("Queue convenience constructor failed: expected 'convenience', got %v", val)
	}

	hashSet := collections.NewHashSetWithCapacity[int](20)
	hashSet.Add(100)

	if !hashSet.Contains(100) {
		t.Error("HashSet convenience constructor failed: value not found")
	}

	concurrentSet := collections.NewConcurrentHashSetWithCapacity[string](15)
	concurrentSet.Add("concurrent")

	if !concurrentSet.Contains("concurrent") {
		t.Error("ConcurrentHashSet convenience constructor failed: value not found")
	}
}

func TestTypeInferenceHelpers(t *testing.T) {
	// Test type inference helpers - types inferred from values!
	stack := collections.StackOf(1, 2, 3, 4, 5)

	if val, ok := stack.Pop(); !ok || val != 5 {
		t.Errorf("StackOf failed: expected 5, got %v", val)
	}

	queue := collections.QueueOf("first", "second", "third")

	if val, ok := queue.Dequeue(); !ok || val != "first" {
		t.Errorf("QueueOf failed: expected 'first', got %v", val)
	}

	hashSet := collections.HashSetOf(10, 20, 30)

	if !hashSet.Contains(20) || hashSet.Size() != 3 {
		t.Error("HashSetOf failed: values not properly added")
	}

	concurrentSet := collections.ConcurrentHashSetOf("a", "b", "c")

	if !concurrentSet.Contains("b") || concurrentSet.Size() != 3 {
		t.Error("ConcurrentHashSetOf failed: values not properly added")
	}
}

func TestFluentBuilders(t *testing.T) {
	// Test fluent builders - maximum expressiveness!
	stack := collections.NewStackBuilder[int]().
		WithCapacity(10).
		WithValues(1, 2, 3).
		Build()

	if val, ok := stack.Pop(); !ok || val != 3 {
		t.Errorf("Stack builder failed: expected 3, got %v", val)
	}

	queue := collections.NewQueueBuilder[string]().
		WithCapacity(5).
		WithValues("alpha", "beta", "gamma").
		Build()

	if val, ok := queue.Dequeue(); !ok || val != "alpha" {
		t.Errorf("Queue builder failed: expected 'alpha', got %v", val)
	}

	// Test builder without initial values
	emptyStack := collections.NewStackBuilder[float64]().
		WithCapacity(100).
		Build()

	if !emptyStack.IsEmpty() {
		t.Error("Empty stack builder should create empty stack")
	}
}
