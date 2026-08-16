# Go Utility Packages Design

## Context

The current repository, `github.com/yuvv/candy`, is a Go 1.18 module. It already contains Java-inspired packages such as `collection`, `iter`, `stream`, `optional`, `function`, and `lang`, but several areas are incomplete or experimental. For example, parts of `collection` and `stream` still contain unimplemented behavior, and the current `optional` implementation relies on reflection-heavy zero-value checks.

The reference repository, `/home/yuweiwei.919/go/src/code.byted.org/ad/brand-starter/lang-starter`, is also a Go 1.18 utility module. Its most useful patterns are small independent utility packages: slice helpers, optional values, cache implementations, concurrent task helpers, string helpers, and lightweight type utilities.

The user approved an "utility-first" direction: add stable, practical Go utility packages inspired by the reference library while avoiding a destructive rewrite of the current unfinished Java-style collection/stream APIs.

## Goals

1. Add practical, independent utility packages that can be used immediately in normal Go code.
2. Keep compatibility risk low by not deleting or replacing existing `collection`, `iter`, or `stream` APIs in this first pass.
3. Prefer clear Go-style APIs over direct line-by-line copying from the reference library.
4. Keep the module compatible with Go 1.18.
5. Cover each new package with focused unit tests.

## Non-Goals

1. Do not complete the whole existing `stream` implementation in this pass.
2. Do not redesign the entire Java-style collection hierarchy in this pass.
3. Do not introduce external dependencies unless they are strictly necessary; the initial implementation should use only the standard library.
4. Do not change the module path.

## Recommended Architecture

Add new packages beside the existing packages:

```text
arrays/
  slice.go
  slice_test.go

optional/
  optional.go
  optional_test.go
  ptr.go
  ptr_test.go

cache/
  cache.go
  lru.go
  lru_test.go

concurrent/
  task_group.go
  batch.go
  task_group_test.go
  batch_test.go

strs/
  builder.go
  joiner.go
  builder_test.go
  joiner_test.go

types/
  bitflag.go
  bitflag_test.go
```

Existing packages remain in place. The new packages should not depend on `collection`, `stream`, or `iter`. That keeps the first pass easy to test and prevents unfinished older APIs from leaking into the new utility surface.

## Package Designs

### arrays

Purpose: provide common generic slice operations for Go 1.18 users without requiring Go 1.21's `slices` package.

Initial functions:

- `Filter[T any](items []T, keep func(T) bool) []T`: return a new slice containing matching items.
- `FilterInPlace[T any](items []T, keep func(T) bool) []T`: compact matching items into the input backing array.
- `Reject[T any](items []T, reject func(T) bool) []T`: inverse of `Filter`.
- `RejectInPlace[T any](items []T, reject func(T) bool) []T`.
- `RemoveZero[T comparable](items []T) []T`: return a new slice without zero values.
- `RemoveZeroInPlace[T comparable](items []T) []T`.
- `Map[T any, R any](items []T, mapper func(T) R) []R`.
- `FlatMap[T any, R any](items []T, mapper func(T) []R) []R`.
- `Reduce[T any, R any](items []T, initial R, reducer func(R, T) R) R`.
- `Contains[T comparable](items []T, target T) bool`.
- `Unique[T comparable](items []T) []T`: preserve first-seen order.
- `Chunk[T any](items []T, size int) [][]T`: panic when `size <= 0` so misuse is explicit.
- `GroupBy[T any, K comparable](items []T, key func(T) K) map[K][]T`.
- `ToMap[T any, K comparable, V any](items []T, key func(T) K, value func(T) V) map[K]V`.

Nil handling: functions that return slices should preserve nil where it matters. For example, filtering a nil slice returns nil; mapping a nil slice returns nil. Functions that build maps return an empty map for empty but non-nil input and an empty map for nil input, because a nil map is less convenient for callers.

### optional

Purpose: provide a simpler Optional implementation with explicit presence semantics.

Current `optional.Optional[T]` uses a value plus reflection-based nil/zero checks. The new design should use:

```go
type Optional[T any] struct {
    value   T
    present bool
}
```

Initial API:

- `Empty[T any]() Optional[T]`
- `Of[T any](value T) Optional[T]`
- `OfPtr[T any](value *T) Optional[T]`: empty when pointer is nil, otherwise stores dereferenced value.
- `IsPresent() bool`
- `IsEmpty() bool`
- `Get() T`: panic with `optional: no value present` when empty.
- `OrElse(defaultValue T) T`
- `OrElseGet(func() T) T`
- `IfPresent(func(T))`
- `Map[T, R]` cannot be a method with a new type parameter in Go 1.18, so expose it as package function `Map[T any, R any](Optional[T], func(T) R) Optional[R]`.
- `FlatMap[T, R]` as package function.

Compatibility note: this changes `Empty`, `Of`, and related functions from pointer-returning to value-returning APIs if implemented directly in the existing package. Because the current branch is unfinished and the user allowed design freedom, this compatibility break is acceptable. Tests and examples should document the new API.

### cache

Purpose: provide cache interfaces and a concurrent-safe LRU cache.

Initial API:

```go
type Cache[K comparable, V any] interface {
    Get(K) (V, bool)
    Put(K, V)
    Remove(K) bool
    Clear()
    Len() int
}
```

`NewLRU[K comparable, V any](capacity int) *LRU[K,V]` creates a cache. It should panic when `capacity <= 0`. `Get` promotes entries to most recently used. `Put` updates existing entries and evicts the least recently used entry when capacity is exceeded. Methods are protected by a mutex.

### concurrent

Purpose: simplify running context-aware tasks over slices.

Initial API:

- `TaskGroup[OUT any]` runs functions of shape `func(context.Context) (OUT, error)`.
- `NewTaskGroup[OUT any](ctx context.Context) *TaskGroup[OUT]`.
- `Go(func(context.Context) (OUT, error))` starts one task.
- `Wait() ([]OUT, error)` waits for all tasks. On first error, cancel the derived context and return that error after all goroutines finish.
- `AwaitTask[IN any, OUT any](ctx context.Context, fn func(context.Context, IN) (OUT, error), inputs []IN) ([]OUT, error)`.
- `AwaitBatchTask[IN any, OUT any](ctx context.Context, fn func(context.Context, []IN) ([]OUT, error), inputs []IN, batchSize int) ([]OUT, error)`.
- `AwaitBatchReturnMapTask[IN comparable, OUT any](ctx context.Context, fn func(context.Context, []IN) (map[IN]OUT, error), inputs []IN, batchSize int) (map[IN]OUT, error)`.

Ordering: `AwaitTask` should preserve input order. Batch helpers should preserve batch order when flattening slice results. Map results have normal map ordering.

Batch size: if `batchSize <= 0`, treat it as 1, matching the reference behavior.

### strs

Purpose: string helper utilities without conflicting with the standard `strings` package name.

Initial API:

- `Builder`: small wrapper over `strings.Builder` with chainable `Append`, `AppendString`, `AppendByte`, `String`, and `Len` methods.
- `JoinSlice[T any](items []T, adapt func(T) string, prefix, delimiter, suffix string) string`.
- `Joiner[T any]` with `NewJoiner(delimiter string, adapt func(T) string)`, `WithPrefix`, `WithSuffix`, and `Join`.
- Convenience constructors for `int`, `int32`, `int64`, `uint`, `uint64`, and `string` joiners.

### types

Purpose: small type utilities.

Initial API:

- `BitFlag[T ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64]`.
- `Has(flag T) bool`.
- `Set(flag T)`.
- `Clear(flag T)`.
- `Toggle(flag T)`.
- `Value() T`.

This package should start with bit flags only. Bitmap can be added later if needed.

## Error Handling

- Invalid constructor parameters, such as LRU capacity `<= 0` or chunk size `<= 0`, should panic. These are programmer errors rather than runtime data errors.
- Concurrent task functions should return errors normally. The task group should cancel sibling tasks after the first observed error.
- Optional `Get` should panic on empty, matching common Optional semantics and making misuse obvious.

## Testing Strategy

Use test-first implementation for each package.

Required tests:

- `arrays`: nil preservation, filter/reject behavior, in-place compaction, map/flatmap/reduce, unique order, chunk boundaries, invalid chunk size panic, group and map conversion.
- `optional`: empty/present state, `Get` panic, default suppliers are lazy, `IfPresent`, `Map`, `FlatMap`, pointer optional behavior.
- `cache`: put/get, update, recency promotion, eviction order, remove return value, clear, invalid capacity panic.
- `concurrent`: successful task collection, input order preservation, batch size behavior, error propagation and cancellation.
- `strs`: builder chaining, join with prefix/suffix/delimiter, numeric joiners.
- `types`: set/clear/toggle/has/value behavior.

Run `go test ./...` before completion.

## Compatibility and Migration Notes

This design intentionally leaves older packages in place except for `optional`, where replacing the existing reflection-based implementation is valuable and aligned with the user-approved freedom to redesign unfinished code. If downstream users depend on the current pointer-returning `optional` API, they will need to adjust. Since this is a development branch with incomplete behavior, the cleaner API is preferred.

## Implementation Order

1. Add `arrays` package with tests.
2. Replace `optional` with explicit presence semantics and tests.
3. Add `cache` LRU with tests.
4. Add `strs` helpers with tests.
5. Add `types` bit flags with tests.
6. Add `concurrent` task helpers with tests.
7. Run full formatting and `go test ./...`.
