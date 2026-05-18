# Overview

The **collections** module provides small, focused data structures built on Go 1.18+ generics. Each type lives in its own package with a consistent optional-capacity pattern.

## Included types

| Type | Package | Thread-safe | Notes |
|------|---------|-------------|-------|
| `HashSet[T]` | `hashset` | No | `map[T]struct{}` backed set |
| `ConcurrentHashSet[T]` | `concurrenthashset` | Yes | `sync.RWMutex` |
| `Queue[T]` | `queue` | No | FIFO with shifting buffer |
| `Stack[T]` | `stack` | No | LIFO stack |

Root package `collections` re-exports constructors (`NewHashSet`, `NewQueue`, `NewStack`, builders) for convenience.

## Design goals

- **Type safety** — no `interface{}` or reflection in hot paths
- **Minimal API** — only operations needed for processor pipelines and in-memory caches
- **Optional preallocation** — `WithCapacity` / `New*WithCapacity` to reduce reallocations
- **Predictable complexity** — documented behavior per method in package godoc

## When to use

- In-process deduplication and membership (`HashSet`)
- Work queues inside a single goroutine or guarded by your own mutex (`Queue`, `Stack`)
- Concurrent caches where `sync.Map` is too loose (`ConcurrentHashSet`)

For persistent or distributed structures, use a database or `github.com/fgrzl/kv` instead.
