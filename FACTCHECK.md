# Fact-Check & Test Report

Full audit of every program in this repository: correctness review, bug fixes,
compile repairs, and a test suite. Final state: **all 27 packages build, pass
`go vet`, and pass `go test ./...`** (79 test files).

## 1. Build foundation (was impossible to compile before)

- Added root `go.mod` (`module github.com/ajit-ai/Golang`) and `use .` in `go.work`.
- Removed `*.mod` / `*.work` from `.gitignore` so the module files are tracked.
- Added missing dependency `github.com/go-sql-driver/mysql`.
- Resolved ~60 duplicate `func main()` declarations across 20 packages.
  Every directory now has exactly one demo-runner `main()` that calls the
  per-file `<Topic>Main()` entry points.
- De-collided package-level identifiers:
  - `Node` → `SinglyNode` (Ch03), `DoublyNode` (Ch03), `UnorderedNode` (Ch06)
  - `TreeNode` → `AVLTreeNode` (Ch04)
  - `Queue` → `SyncQueue` (Ch03)
  - `ReferenceCounter` → `StackReferenceCounter` (Ch10 stack_garbage_collection.go)
  - `addOne` → `addOnePointer` (Ch10 stack_memory_pointer.go)
  - `hash` func → `stringHash` (Ch08; collided with the `"hash"` import)
  - Ch02 CRM API prefixed (`CrmCustomer`, `GetCrmConnection`, …) to coexist with
    database_operations.go; `Home` handler → `WebFormsHome`
  - Ch09 map-based graph → `ExampleSocialGraph` / `NewExampleSocialGraph`
- Fixed book-style shadowed variables (`var Node *Node`, `var CrmCustomer CrmCustomer`)
  in linked lists and CRM code.
- Renamed misnamed `selection-sort._test.go` → `selection-sort_test.go`;
  relocated `basic/Chapter09/test/main_test.go` → `basic/Chapter09/constructors_test.go`.

Blocking demos are intentionally NOT wired into auto-runners (documented in each runner):
`SyncQueueMain` (ends in `select {}`), `CrmAppMain`/`WebFormsMain` (serve HTTP on :8000),
`DatabaseOperationsMain` (needs live MySQL), `demo.AdvancedMain` (starts an HTTP server).

## 2. Algorithm bugs found and fixed

| # | Location | Bug | Fix |
|---|----------|-----|-----|
| 1 | `algorithms/merge-sort/merge-sort.go` | Variants 1–5 dropped the remaining right half when the left half exhausted first — `MergeSortInt([1,3,2])` returned `[1]` | Append both remainders (as variants 6–10 already did); regression tests added |
| 2 | `basic/Chapter01/backtracking.go` | Wrote `combinations[m]` with no bounds check; `[19]int` buffer overflowed for large targets (index-out-of-range panic) | Guard after the equality-count, before the write loop |
| 3 | `basic/Chapter03/stack.go` | `Pop()` kept the last element in the slice, so `elementCount` never reached 0 — the final element could be popped infinitely | Truncate with `elements[:0]` |
| 4 | `basic/Chapter03/sync_queue.go` | Demo deadlocks by design (`select {}`) | Excluded from runner; deterministic handshake test added |
| 5 | `demo/basic.go` `quickSort` | Empty `else` branch discarded every element greater than the pivot — `[5,2,9,1,5,6]` sorted to `[1,2,5]` | Restore right partition append |
| 6 | `demo/basic.go` `sieveOfEratosthenes(arr)` | Ignored its input entirely; always returned primes ≤ 100 | Sieve up to max(input); filter the actual input |
| 7 | `demo/basic.go` `numberToBinary` | `string(n%2)` converts int→rune (control chars), not digits | `strconv.Itoa(n%2)` |
| 8 | `arrays/duplicate/duplicate.go` `findDuplicateseleven` | Mutex locked but never unlocked → deadlock on second iteration | Unlock at end of each iteration |
| 9 | `arrays/duplicate/duplicate.go` `findDuplicatesten` | `wg.Add(1)` with two `Done()` calls → negative WaitGroup counter panic | `wg.Add(2)`, both goroutines defer `Done()` |
| 10 | `arrays/duplicate/duplicate.go` `findDuplicatesseven` | "Recursive" variant recursed into the map-based `findDuplicates`, not itself | Self-recursive call |
| 11 | `arrays/duplicate/duplicate.go` `findDuplicatesfour` | Bool slice sized `len(arr)` but indexed by element value → out-of-range for large values | Size by max value |
| 12 | `basic/Chapter04/binary_search_tree.go` + `basic/Chapter07/binarysearchtree.go` | `RemoveNode` discarded `removeNode`'s return value — removing the ROOT left `rootNode` dangling and searchable | `tree.rootNode = removeNode(tree.rootNode, key)`; root-removal regression tests |
| 13 | `basic/Chapter06/sort_multi_keys.go` | `Swap` copied j→i then new-i→j: a no-op that duplicated values, so multi-key sort did nothing | Tuple swap |
| 14 | `basic/Chapter08/hash_string.go` | `stringHash` Atoi-parsed the whole string, ignored the error, always returned 0 | Polynomial rolling hash (constant 42) |
| 15 | `basic/Chapter05/twodmatrix.go` `inverse` | Never compiled: scalar float where `[][]float64` was required; int/float division mismatches | Proper 2D inverse; singular matrix returns nil |
| 16 | `constant/constantex.go` | Unused variable `m` — package never compiled | Use the created map |

## 3. Code that never compiled and had to be implemented

- `basic/Chapter10/mark.go`, `sweep.go`, `generation_collect.go`: book pseudo-code
  using invalid syntax (`map[root] = true`) and undefined symbols. Implemented a
  working GC simulation core (`object`, mark bits, live heap, `Release`) with
  correct mark/sweep/generational semantics, plus unit tests.
- `basic/Chapter10/weighted_reference.go`: no package clause, undefined
  `GetReferences(root)`, duplicate type name. Rewritten self-contained.

## 4. Latent runtime landmines defused

- `basic/Chapter09/social_graph.go`: demo main was commented out because it
  panicked — `NewSocialGraph(4)` then `AddLink(2, 4, …)` indexes slot 4 of a
  length-4 slice. Uncommented with size fixed to 5.
- `basic/Chapter02/crm_app.go`: package-level `template.Must(ParseGlob("templates/*"))`
  panicked whenever CWD wasn't the package directory (including under `go run ./...`).
  Now fault-tolerant via `mustParseTemplates()`.
- `basic/Chapter09/social_graph_example.go`: `%d` verbs for string-typed `Name`
  values (failed `go vet`, which `go test` runs). Changed to `%s`.
- `basic/Chapter08/interpolation_search.go`: unreachable trailing `return` (vet).
- `basic/Chapter08` insertion/merge sort demos: redundant `\n` and `%d` inside
  `fmt.Println` (vet printf checks).
- `basic/Chapter01/backtracking_test.go` exposed bug #2 above.

## 5. Known quirks documented but intentionally preserved

- `basic/Chapter10/reference_counting.go` `Subtract()`: decrementing below zero
  wraps `num` to 2^32−1; the "removed" counter only increments on the exact 1→0
  transition. Matches the book's snippet; tests pin the real behavior.
- `functions/functionsuse.go` `divide(a, 0)` returns 0 instead of erroring
  (guarded by design; tested).
- `arrays/duplicate/duplicate.go` main labels #6 "binary search" but reuses the
  sorting variant (#5); cosmetic only.
- `demo/basic.go` `atoi` and other helpers mirror book examples whose edge-case
  semantics differ from the standard library.

## 6. Test suite

79 `_test.go` files covering all 27 packages:

- Regression tests for every numbered bug above.
- Table-driven logic tests for data structures (BST incl. root removal, AVL
  ordering, stacks, circular queue wraparound/overflow panic, sets, singly and
  doubly linked lists, sparse matrix shape, dictionary CRUD, tree set).
- Sort/search algorithm tests (bubble, insertion, merge incl. left-exhausted
  path, selection, quick, shell, linear/interpolation search, binary search).
- Concurrency tests with bounded timeouts (sync queue handshake, goroutine /
  WaitGroup / mutex duplicate finders).
- Smoke tests invoking every demo entry point (must not panic).
