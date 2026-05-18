# Contributing

Thanks for contributing to collections.

## Setup

1. Fork and clone the repository.
2. `go mod download`
3. `go test ./...`

## Pull requests

- Run `go fmt ./...` and `go vet ./...`.
- Add table-driven tests for new collection types or methods.
- Update `docs/` for API changes.
- Keep the library small and focused—avoid scope creep.

## Changelog

Note changes under `## [Unreleased]` in [CHANGELOG.md](CHANGELOG.md).
