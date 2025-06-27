[![CI](https://github.com/fgrzl/collections/actions/workflows/ci.yml/badge.svg)](https://github.com/fgrzl/collections/actions/workflows/ci.yml)
[![Dependabot Updates](https://github.com/fgrzl/collections/actions/workflows/dependabot/dependabot-updates/badge.svg)](https://github.com/fgrzl/collections/actions/workflows/dependabot/dependabot-updates)

# Collections

A fast, generic, and type-safe collections library for Go.

## Features

- Type-safe, generic implementations using Go 1.18+ generics
- Minimal memory overhead and optimized performance
- Modular design with consistent APIs and optional capacity preallocation
- Thread-safe variants where appropriate

## Included Collections

### `hashset.HashSet[T]`
A simple, non-thread-safe hash set.

- Backed by `map[T]struct{}`
- Operations: `Add`, `Remove`, `Contains`, `IsEmpty`, `Size`, `Clear`, `ToSlice`, `ForEach`
- Optional: `WithCapacity`

### `concurrenthashset.ConcurrentHashSet[T]`
Thread-safe version of `HashSet`.

- Uses `sync.RWMutex` for concurrent access
- Same API as `HashSet`
- Optional: `WithCapacity`

### `queue.Queue[T]`
An optimized FIFO queue with internal shifting and reallocation.

- Fast `Enqueue`, `Dequeue`, and `Length` operations
- Avoids unnecessary allocations with shifting and shrinking
- Optional: `WithCapacity`

### `stack.Stack[T]`
Simple LIFO stack.

- Operations: `Push`, `Pop`, `Peek`, `Length`, `IsEmpty`, `Reset`
- Optional: `WithCapacity`

## Installation

```bash
go get github.com/fgrzl/collections
```