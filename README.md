### Golang — a hands-on Go learning repository

A structured collection of runnable Go example programs, from language basics
to concurrency patterns, each with unit tests. Every directory is a
self-contained `package main` demo (one `main()` per directory that calls the
per-file `<Topic>Main()` entry points), except `pkg/`, which is a real
importable library.

## Repository layout

| Folder | Contents |
|---|---|
| `basics/` | variables, constants, expressions, for, if/else, switch, functions |
| `collections/` | arrays, slices, maps (incl. 100-example deep dives) |
| `language/` | generics, interfaces, structs & embedding, modern error handling |
| `algorithms/` | bubble, insertion, merge, selection sort with tests |
| `basic/Chapter01–10` | book-style chapters: patterns, complexity, data structures, matrices, sorting, graphs, memory & GC simulation |
| `concurrency/` | worker pools, pipelines, fan-in/fan-out, semaphores, graceful shutdown (race-detector clean) |
| `stdlib/` | JSON, file I/O & CSV, HTTP JSON API (Go 1.22 routing), CLI flags, context |
| `demo/` | large mixed showcase incl. 80+ classic problems, benchmarks, fuzzing |
| `testing/` | minimal introduction to Go tests |
| `pkg/gostack/` | importable generic stack library |
| `examples/` | programs that consume `pkg/` libraries |

## Quick start

Requires Go 1.22+ (`go.mod` / `go.work` included).

```bash
# run any demo (from the repository root)
go run ./basics/functions
go run ./language/generics
go run ./concurrency

# run every test
go test ./...

# race-detector pass over the concurrency examples
go test -race ./concurrency

# benchmarks: strings.Builder vs += concatenation
go test ./demo -bench . -run XXX

# fuzzing (discovers inputs automatically)
go test ./demo -run FuzzIsPalindrome -fuzz FuzzIsPalindrome -fuzztime 30s

# library usage example
go run ./examples/gostack-demo
```

## Notes

- Demos that need external resources are excluded from auto-runners and must be
  started manually: `stdlib/httpapi` (serves :8080), `basic/Chapter02`
  CRM/webform servers (MySQL + :8000), `basic/Chapter03` SyncQueueMain
  (endless demo).
- `FACTCHECK.md` documents the full audit of this code base: bugs found,
  fixes applied, and known quirks kept intentionally.
- CI (`.github/workflows/go.yml`) runs gofmt check, vet, build and
  `go test -race` on pushes and PRs to `main` and `develop`.
