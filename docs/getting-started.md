# Getting started

## Install

```bash
go get github.com/fgrzl/collections
```

## Hash set

```go
import "github.com/fgrzl/collections/hashset"

set := hashset.NewHashSet[string](hashset.WithCapacity(64))
set.Add("alpha", "beta")
if set.Contains("alpha") {
    // ...
}
```

Or via the root package:

```go
import "github.com/fgrzl/collections"

set := collections.NewHashSetWithCapacity[string](64)
```

## Queue (FIFO)

```go
import "github.com/fgrzl/collections/queue"

q := queue.NewQueue[int]()
q.Enqueue(1, 2, 3)
v, ok := q.Dequeue() // 1, true
```

## Stack (LIFO)

```go
import "github.com/fgrzl/collections/stack"

s := stack.NewStack[string]()
s.Push("bottom", "top")
v, ok := s.Pop() // "top", true
```

## Concurrent hash set

```go
import "github.com/fgrzl/collections/concurrenthashset"

set := concurrenthashset.NewConcurrentHashSet[int]()
set.Add(42)
```

## Tests

```bash
go test ./...
```
