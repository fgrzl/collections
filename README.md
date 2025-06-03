[![ci](https://github.com/fgrzl/collections/actions/workflows/ci.yml/badge.svg)](https://github.com/fgrzl/collections/actions/workflows/ci.yml)
[![Dependabot Updates](https://github.com/fgrzl/collections/actions/workflows/dependabot/dependabot-updates/badge.svg)](https://github.com/fgrzl/collections/actions/workflows/dependabot/dependabot-updates)

# Collections

## HashSet

A simple, type-safe, generic HashSet implementation in Go.

Built on top of Go's efficient map[T]struct{} idiom, HashSet provides fast membership testing, insertion, deletion, and iteration for any comparable type.

Features
Generic over any comparable type

Constant time O(1) operations for Add, Remove, and Contains

Simple API: Add, Remove, Contains, Size, IsEmpty, Clear, ToSlice, and ForEach

Zero unnecessary memory allocations

## ConcurrentHashSet

A simple, type-safe, generic ConcurrentHashSet implementation in Go.

Built on top of Go's efficient map[T]struct{} idiom, ConcurrentHashSet provides fast membership testing, insertion, deletion, and iteration for any comparable type.

Features
Generic over any comparable type

Constant time O(1) operations for Add, Remove, and Contains

Simple API: Add, Remove, Contains, Size, IsEmpty, Clear, ToSlice, and ForEach

Zero unnecessary memory allocations

