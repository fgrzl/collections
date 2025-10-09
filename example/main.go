// Example demonstrating the "pit of success" improvements

package main

import (
	"fmt"

	"github.com/fgrzl/collections"
)

func main() {
	// BEFORE: Verbose and error-prone
	// stack := collections.NewStack[int](collections.StackWithCapacity[int](10))
	// hashSet := collections.NewHashSet[string](collections.HashSetWithCapacity[string](20))

	// AFTER: Multiple ergonomic options!

	// 1. Convenience constructors - clean and simple
	_ = collections.NewStackWithCapacity[int](10)
	_ = collections.NewQueueWithCapacity[string](5)

	// 2. Type inference - types inferred from values!
	numbers := collections.StackOf(1, 2, 3, 4, 5)
	words := collections.QueueOf("hello", "world", "collections")
	tags := collections.HashSetOf("go", "generics", "collections", "ergonomic")

	// 3. Fluent builders - maximum expressiveness
	processQueue := collections.NewQueueBuilder[string]().
		WithCapacity(100).
		WithValues("task1", "task2", "task3").
		Build()

	// 4. Mix and match as needed
	cache := collections.NewConcurrentHashSetWithCapacity[string](1000)

	// Usage examples
	fmt.Println("Stack length:", numbers.Length())
	fmt.Println("Queue length:", words.Length())
	fmt.Println("HashSet contains 'go':", tags.Contains("go"))
	fmt.Println("Process queue has tasks:", !processQueue.IsEmpty())

	cache.Add("cached-item")
	fmt.Println("Cache size:", cache.Size())
}
