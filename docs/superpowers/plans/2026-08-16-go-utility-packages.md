# Go Utility Packages Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add practical Go 1.18 utility packages inspired by lang-starter: arrays, optional, cache, strs, types, and concurrent.

**Architecture:** New utility packages are independent and do not depend on the existing unfinished collection/stream hierarchy. The existing optional package is replaced with explicit presence semantics while preserving the exported `Optional[T]` type name used by stream interfaces.

**Tech Stack:** Go 1.18, standard library only, `testing` package, generics.

---

## File Structure

- Create `arrays/slice.go`: generic slice helpers.
- Create `arrays/slice_test.go`: tests for all slice helpers.
- Replace `optional/optional.go`: explicit-present Optional implementation.
- Create `optional/ptr.go`: pointer Optional constructor.
- Replace `optional/optional_test.go`: tests for the new Optional API.
- Create `optional/ptr_test.go`: pointer Optional tests.
- Create `cache/cache.go`: generic cache interface.
- Create `cache/lru.go`: concurrent-safe LRU implementation.
- Create `cache/lru_test.go`: LRU behavior tests.
- Create `strs/builder.go`: chainable string builder wrapper.
- Create `strs/joiner.go`: generic join helpers.
- Create `strs/builder_test.go` and `strs/joiner_test.go`.
- Create `types/bitflag.go` and `types/bitflag_test.go`.
- Create `concurrent/task_group.go`: context-aware task group.
- Create `concurrent/batch.go`: batch helpers.
- Create `concurrent/task_group_test.go` and `concurrent/batch_test.go`.
- Modify `README.md`: concise package overview.

---

### Task 1: Add arrays slice helpers

**Files:**
- Create: `arrays/slice_test.go`
- Create: `arrays/slice.go`

- [ ] **Step 1: Write the failing arrays tests**

Create `arrays/slice_test.go`:

```go
package arrays

import (
	"reflect"
	"testing"
)

func TestFilterRejectAndNilPreservation(t *testing.T) {
	var nilInts []int
	if got := Filter(nilInts, func(v int) bool { return v > 0 }); got != nil {
		t.Fatalf("Filter(nil) = %#v, want nil", got)
	}

	items := []int{1, 2, 3, 4}
	if got := Filter(items, func(v int) bool { return v%2 == 0 }); !reflect.DeepEqual(got, []int{2, 4}) {
		t.Fatalf("Filter even = %#v", got)
	}
	if got := Reject(items, func(v int) bool { return v%2 == 0 }); !reflect.DeepEqual(got, []int{1, 3}) {
		t.Fatalf("Reject even = %#v", got)
	}
}

func TestFilterRejectInPlace(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	got := FilterInPlace(items, func(v int) bool { return v%2 == 1 })
	if !reflect.DeepEqual(got, []int{1, 3, 5}) {
		t.Fatalf("FilterInPlace odd = %#v", got)
	}
	if &got[0] != &items[0] {
		t.Fatalf("FilterInPlace should reuse the input backing array")
	}

	items = []int{1, 2, 3, 4, 5}
	got = RejectInPlace(items, func(v int) bool { return v < 3 })
	if !reflect.DeepEqual(got, []int{3, 4, 5}) {
		t.Fatalf("RejectInPlace <3 = %#v", got)
	}
}

func TestRemoveZero(t *testing.T) {
	items := []string{"", "a", "", "b"}
	if got := RemoveZero(items); !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Fatalf("RemoveZero = %#v", got)
	}

	items = []string{"", "a", "", "b"}
	got := RemoveZeroInPlace(items)
	if !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Fatalf("RemoveZeroInPlace = %#v", got)
	}
}

func TestMapFlatMapReduce(t *testing.T) {
	items := []int{1, 2, 3}
	if got := Map(items, func(v int) string { return string(rune('a' + v - 1)) }); !reflect.DeepEqual(got, []string{"a", "b", "c"}) {
		t.Fatalf("Map = %#v", got)
	}
	if got := FlatMap(items, func(v int) []int { return []int{v, v * 10} }); !reflect.DeepEqual(got, []int{1, 10, 2, 20, 3, 30}) {
		t.Fatalf("FlatMap = %#v", got)
	}
	if got := Reduce(items, 10, func(sum int, v int) int { return sum + v }); got != 16 {
		t.Fatalf("Reduce = %d", got)
	}
}

func TestContainsUniqueChunk(t *testing.T) {
	items := []int{1, 2, 1, 3, 2}
	if !Contains(items, 3) || Contains(items, 4) {
		t.Fatalf("Contains returned wrong result")
	}
	if got := Unique(items); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Fatalf("Unique = %#v", got)
	}
	if got := Chunk([]int{1, 2, 3, 4, 5}, 2); !reflect.DeepEqual(got, [][]int{{1, 2}, {3, 4}, {5}}) {
		t.Fatalf("Chunk = %#v", got)
	}
}

func TestChunkPanicsForInvalidSize(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatalf("Chunk should panic for size <= 0")
		}
	}()
	_ = Chunk([]int{1}, 0)
}

func TestGroupByAndToMap(t *testing.T) {
	items := []string{"go", "hi", "tool"}
	groups := GroupBy(items, func(v string) int { return len(v) })
	if !reflect.DeepEqual(groups[2], []string{"go", "hi"}) || !reflect.DeepEqual(groups[4], []string{"tool"}) {
		t.Fatalf("GroupBy = %#v", groups)
	}

	mapped := ToMap(items, func(v string) int { return len(v) }, func(v string) string { return v + "!" })
	if !reflect.DeepEqual(mapped, map[int]string{2: "hi!", 4: "tool!"}) {
		t.Fatalf("ToMap = %#v", mapped)
	}
}
```

- [ ] **Step 2: Run arrays tests and verify RED**

Run:

```bash
go test ./arrays
```

Expected: FAIL because `arrays` package functions do not exist.

- [ ] **Step 3: Implement arrays helpers**

Create `arrays/slice.go`:

```go
package arrays

func Filter[T any](items []T, keep func(T) bool) []T {
	if items == nil {
		return nil
	}
	out := make([]T, 0, len(items))
	for _, item := range items {
		if keep(item) {
			out = append(out, item)
		}
	}
	return out
}

func FilterInPlace[T any](items []T, keep func(T) bool) []T {
	j := 0
	for _, item := range items {
		if keep(item) {
			items[j] = item
			j++
		}
	}
	return items[:j]
}

func Reject[T any](items []T, reject func(T) bool) []T {
	return Filter(items, func(item T) bool { return !reject(item) })
}

func RejectInPlace[T any](items []T, reject func(T) bool) []T {
	return FilterInPlace(items, func(item T) bool { return !reject(item) })
}

func RemoveZero[T comparable](items []T) []T {
	var zero T
	return Reject(items, func(item T) bool { return item == zero })
}

func RemoveZeroInPlace[T comparable](items []T) []T {
	var zero T
	return RejectInPlace(items, func(item T) bool { return item == zero })
}

func Map[T any, R any](items []T, mapper func(T) R) []R {
	if items == nil {
		return nil
	}
	out := make([]R, 0, len(items))
	for _, item := range items {
		out = append(out, mapper(item))
	}
	return out
}

func FlatMap[T any, R any](items []T, mapper func(T) []R) []R {
	if items == nil {
		return nil
	}
	out := make([]R, 0, len(items))
	for _, item := range items {
		out = append(out, mapper(item)...)
	}
	return out
}

func Reduce[T any, R any](items []T, initial R, reducer func(R, T) R) R {
	result := initial
	for _, item := range items {
		result = reducer(result, item)
	}
	return result
}

func Contains[T comparable](items []T, target T) bool {
	for _, item := range items {
		if item == target {
			return true
		}
	}
	return false
}

func Unique[T comparable](items []T) []T {
	if items == nil {
		return nil
	}
	seen := make(map[T]struct{}, len(items))
	out := make([]T, 0, len(items))
	for _, item := range items {
		if _, ok := seen[item]; ok {
			continue
		}
		seen[item] = struct{}{}
		out = append(out, item)
	}
	return out
}

func Chunk[T any](items []T, size int) [][]T {
	if size <= 0 {
		panic("arrays: chunk size must be positive")
	}
	if items == nil {
		return nil
	}
	out := make([][]T, 0, (len(items)+size-1)/size)
	for start := 0; start < len(items); start += size {
		end := start + size
		if end > len(items) {
			end = len(items)
		}
		out = append(out, items[start:end])
	}
	return out
}

func GroupBy[T any, K comparable](items []T, key func(T) K) map[K][]T {
	out := make(map[K][]T)
	for _, item := range items {
		k := key(item)
		out[k] = append(out[k], item)
	}
	return out
}

func ToMap[T any, K comparable, V any](items []T, key func(T) K, value func(T) V) map[K]V {
	out := make(map[K]V, len(items))
	for _, item := range items {
		out[key(item)] = value(item)
	}
	return out
}
```

- [ ] **Step 4: Run arrays tests and verify GREEN**

Run:

```bash
go test ./arrays
```

Expected: PASS.

- [ ] **Step 5: Commit arrays package**

```bash
git add arrays/slice.go arrays/slice_test.go
git commit -m "feat: add generic slice helpers"
```

---

### Task 2: Replace optional with explicit presence semantics

**Files:**
- Replace: `optional/optional_test.go`
- Replace: `optional/optional.go`
- Create: `optional/ptr.go`
- Create: `optional/ptr_test.go`

- [ ] **Step 1: Replace optional tests with new API tests**

Replace `optional/optional_test.go`:

```go
package optional

import "testing"

func TestOptionalPresenceAndGet(t *testing.T) {
	empty := Empty[int]()
	if empty.IsPresent() || !empty.IsEmpty() {
		t.Fatalf("empty optional presence mismatch")
	}

	present := Of(42)
	if !present.IsPresent() || present.IsEmpty() {
		t.Fatalf("present optional presence mismatch")
	}
	if got := present.Get(); got != 42 {
		t.Fatalf("Get = %d", got)
	}
}

func TestOptionalGetPanicsWhenEmpty(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatalf("Get should panic on empty optional")
		}
	}()
	_ = Empty[string]().Get()
}

func TestOptionalDefaultsAreLazy(t *testing.T) {
	present := Of("value")
	called := false
	if got := present.OrElse("default"); got != "value" {
		t.Fatalf("OrElse = %q", got)
	}
	if got := present.OrElseGet(func() string { called = true; return "default" }); got != "value" {
		t.Fatalf("OrElseGet = %q", got)
	}
	if called {
		t.Fatalf("OrElseGet supplier should not be called for present optional")
	}

	empty := Empty[string]()
	if got := empty.OrElseGet(func() string { called = true; return "default" }); got != "default" || !called {
		t.Fatalf("empty OrElseGet = %q called=%v", got, called)
	}
}

func TestIfPresentMapAndFlatMap(t *testing.T) {
	seen := 0
	Of(10).IfPresent(func(v int) { seen = v })
	if seen != 10 {
		t.Fatalf("IfPresent saw %d", seen)
	}
	Empty[int]().IfPresent(func(v int) { t.Fatalf("IfPresent should not run for empty optional") })

	mapped := Map(Of(10), func(v int) string { return "v" })
	if !mapped.IsPresent() || mapped.Get() != "v" {
		t.Fatalf("Map result = %#v", mapped)
	}
	if Map(Empty[int](), func(v int) string { return "bad" }).IsPresent() {
		t.Fatalf("Map should keep empty optional empty")
	}

	flat := FlatMap(Of(10), func(v int) Optional[int] { return Of(v * 2) })
	if !flat.IsPresent() || flat.Get() != 20 {
		t.Fatalf("FlatMap result = %#v", flat)
	}
}
```

Create `optional/ptr_test.go`:

```go
package optional

import "testing"

func TestOfPtr(t *testing.T) {
	if OfPtr[int](nil).IsPresent() {
		t.Fatalf("OfPtr(nil) should be empty")
	}

	value := 7
	opt := OfPtr(&value)
	if !opt.IsPresent() || opt.Get() != 7 {
		t.Fatalf("OfPtr(&value) = %#v", opt)
	}

	value = 8
	if opt.Get() != 7 {
		t.Fatalf("OfPtr should store dereferenced value, got %d", opt.Get())
	}
}
```

- [ ] **Step 2: Run optional tests and verify RED**

Run:

```bash
go test ./optional
```

Expected: FAIL because the existing Optional returns pointers, lacks `OrElseGet` and package-level `Map`/`FlatMap`, and still has `OfNillable` semantics.

- [ ] **Step 3: Implement new optional package**

Replace `optional/optional.go`:

```go
package optional

const emptyPanicMessage = "optional: no value present"

type Optional[T any] struct {
	value   T
	present bool
}

func Empty[T any]() Optional[T] {
	return Optional[T]{}
}

func Of[T any](value T) Optional[T] {
	return Optional[T]{value: value, present: true}
}

func (o Optional[T]) IsPresent() bool {
	return o.present
}

func (o Optional[T]) IsEmpty() bool {
	return !o.present
}

func (o Optional[T]) Get() T {
	if !o.present {
		panic(emptyPanicMessage)
	}
	return o.value
}

func (o Optional[T]) OrElse(defaultValue T) T {
	if o.present {
		return o.value
	}
	return defaultValue
}

func (o Optional[T]) OrElseGet(supplier func() T) T {
	if o.present {
		return o.value
	}
	return supplier()
}

func (o Optional[T]) IfPresent(consumer func(T)) {
	if o.present {
		consumer(o.value)
	}
}

func Map[T any, R any](o Optional[T], mapper func(T) R) Optional[R] {
	if o.IsEmpty() {
		return Empty[R]()
	}
	return Of(mapper(o.Get()))
}

func FlatMap[T any, R any](o Optional[T], mapper func(T) Optional[R]) Optional[R] {
	if o.IsEmpty() {
		return Empty[R]()
	}
	return mapper(o.Get())
}
```

Create `optional/ptr.go`:

```go
package optional

func OfPtr[T any](value *T) Optional[T] {
	if value == nil {
		return Empty[T]()
	}
	return Of(*value)
}
```

- [ ] **Step 4: Run optional and dependent package tests**

Run:

```bash
go test ./optional ./stream
```

Expected: PASS or stream compile success. If stream fails only because it expected pointer Optional return values in implementations, adjust stream implementation return sites to return value Optional. Do not redesign stream behavior in this task.

- [ ] **Step 5: Commit optional rewrite**

```bash
git add optional/optional.go optional/optional_test.go optional/ptr.go optional/ptr_test.go
git commit -m "feat: simplify optional semantics"
```

---

### Task 3: Add concurrent-safe LRU cache

**Files:**
- Create: `cache/cache.go`
- Create: `cache/lru.go`
- Create: `cache/lru_test.go`

- [ ] **Step 1: Write failing LRU tests**

Create `cache/lru_test.go`:

```go
package cache

import "testing"

func TestLRUPutGetUpdateAndLen(t *testing.T) {
	c := NewLRU[string, int](2)
	c.Put("a", 1)
	c.Put("b", 2)
	c.Put("a", 3)

	if got, ok := c.Get("a"); !ok || got != 3 {
		t.Fatalf("Get(a) = %d,%v", got, ok)
	}
	if c.Len() != 2 {
		t.Fatalf("Len = %d", c.Len())
	}
}

func TestLRUEvictsLeastRecentlyUsed(t *testing.T) {
	c := NewLRU[string, int](2)
	c.Put("a", 1)
	c.Put("b", 2)
	if _, ok := c.Get("a"); !ok {
		t.Fatalf("expected a to be present")
	}
	c.Put("c", 3)

	if _, ok := c.Get("b"); ok {
		t.Fatalf("b should be evicted")
	}
	if got, ok := c.Get("a"); !ok || got != 1 {
		t.Fatalf("a should remain, got %d,%v", got, ok)
	}
	if got, ok := c.Get("c"); !ok || got != 3 {
		t.Fatalf("c should be present, got %d,%v", got, ok)
	}
}

func TestLRURemoveClearAndMissing(t *testing.T) {
	c := NewLRU[string, int](2)
	c.Put("a", 1)
	if !c.Remove("a") {
		t.Fatalf("Remove existing key should return true")
	}
	if c.Remove("missing") {
		t.Fatalf("Remove missing key should return false")
	}
	c.Put("b", 2)
	c.Clear()
	if c.Len() != 0 {
		t.Fatalf("Len after Clear = %d", c.Len())
	}
	if _, ok := c.Get("b"); ok {
		t.Fatalf("b should be absent after Clear")
	}
}

func TestNewLRUPanicsForInvalidCapacity(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatalf("NewLRU should panic for invalid capacity")
		}
	}()
	_ = NewLRU[string, int](0)
}
```

- [ ] **Step 2: Run cache tests and verify RED**

Run:

```bash
go test ./cache
```

Expected: FAIL because `cache` package does not exist.

- [ ] **Step 3: Implement Cache interface and LRU**

Create `cache/cache.go`:

```go
package cache

type Cache[K comparable, V any] interface {
	Get(K) (V, bool)
	Put(K, V)
	Remove(K) bool
	Clear()
	Len() int
}
```

Create `cache/lru.go`:

```go
package cache

import (
	"container/list"
	"sync"
)

type entry[K comparable, V any] struct {
	key   K
	value V
}

type LRU[K comparable, V any] struct {
	mu       sync.Mutex
	capacity int
	items    map[K]*list.Element
	order    *list.List
}

func NewLRU[K comparable, V any](capacity int) *LRU[K, V] {
	if capacity <= 0 {
		panic("cache: LRU capacity must be positive")
	}
	return &LRU[K, V]{
		capacity: capacity,
		items:    make(map[K]*list.Element, capacity),
		order:    list.New(),
	}
}

func (c *LRU[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	elem, ok := c.items[key]
	if !ok {
		var zero V
		return zero, false
	}
	c.order.MoveToFront(elem)
	return elem.Value.(*entry[K, V]).value, true
}

func (c *LRU[K, V]) Put(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, ok := c.items[key]; ok {
		elem.Value.(*entry[K, V]).value = value
		c.order.MoveToFront(elem)
		return
	}

	elem := c.order.PushFront(&entry[K, V]{key: key, value: value})
	c.items[key] = elem
	if len(c.items) > c.capacity {
		c.removeOldestLocked()
	}
}

func (c *LRU[K, V]) Remove(key K) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	elem, ok := c.items[key]
	if !ok {
		return false
	}
	c.removeElementLocked(elem)
	return true
}

func (c *LRU[K, V]) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items = make(map[K]*list.Element, c.capacity)
	c.order.Init()
}

func (c *LRU[K, V]) Len() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.items)
}

func (c *LRU[K, V]) removeOldestLocked() {
	oldest := c.order.Back()
	if oldest != nil {
		c.removeElementLocked(oldest)
	}
}

func (c *LRU[K, V]) removeElementLocked(elem *list.Element) {
	c.order.Remove(elem)
	delete(c.items, elem.Value.(*entry[K, V]).key)
}
```

- [ ] **Step 4: Run cache tests and verify GREEN**

Run:

```bash
go test ./cache
```

Expected: PASS.

- [ ] **Step 5: Commit cache package**

```bash
git add cache/cache.go cache/lru.go cache/lru_test.go
git commit -m "feat: add lru cache"
```

---

### Task 4: Add strs and types helpers

**Files:**
- Create: `strs/builder.go`
- Create: `strs/joiner.go`
- Create: `strs/builder_test.go`
- Create: `strs/joiner_test.go`
- Create: `types/bitflag.go`
- Create: `types/bitflag_test.go`

- [ ] **Step 1: Write failing strs and types tests**

Create `strs/builder_test.go`:

```go
package strs

import "testing"

func TestBuilderChaining(t *testing.T) {
	b := NewBuilder().AppendString("go").AppendByte('-').Append(18)
	if got := b.String(); got != "go-18" {
		t.Fatalf("String = %q", got)
	}
	if b.Len() != len("go-18") {
		t.Fatalf("Len = %d", b.Len())
	}
}
```

Create `strs/joiner_test.go`:

```go
package strs

import "testing"

func TestJoinSliceAndJoiner(t *testing.T) {
	got := JoinSlice([]int{1, 2, 3}, func(v int) string { return NewBuilder().Append(v).String() }, "[", ",", "]")
	if got != "[1,2,3]" {
		t.Fatalf("JoinSlice = %q", got)
	}

	joiner := NewStringJoiner("|").WithPrefix("<").WithSuffix(">")
	if got := joiner.Join([]string{"a", "b"}); got != "<a|b>" {
		t.Fatalf("Joiner.Join = %q", got)
	}
}

func TestNumericJoiners(t *testing.T) {
	if got := NewIntJoiner(",").Join([]int{1, 2}); got != "1,2" {
		t.Fatalf("NewIntJoiner = %q", got)
	}
	if got := NewUint64Joiner(";").Join([]uint64{3, 4}); got != "3;4" {
		t.Fatalf("NewUint64Joiner = %q", got)
	}
}
```

Create `types/bitflag_test.go`:

```go
package types

import "testing"

func TestBitFlag(t *testing.T) {
	var flags BitFlag[uint8]
	flags.Set(1)
	flags.Set(4)
	if !flags.Has(1) || !flags.Has(4) || flags.Has(2) {
		t.Fatalf("unexpected flags after Set: %d", flags.Value())
	}
	flags.Clear(1)
	if flags.Has(1) || flags.Value() != 4 {
		t.Fatalf("unexpected flags after Clear: %d", flags.Value())
	}
	flags.Toggle(2)
	if !flags.Has(2) || flags.Value() != 6 {
		t.Fatalf("unexpected flags after Toggle on: %d", flags.Value())
	}
	flags.Toggle(2)
	if flags.Has(2) || flags.Value() != 4 {
		t.Fatalf("unexpected flags after Toggle off: %d", flags.Value())
	}
}
```

- [ ] **Step 2: Run tests and verify RED**

Run:

```bash
go test ./strs ./types
```

Expected: FAIL because packages do not exist.

- [ ] **Step 3: Implement strs and types**

Create `strs/builder.go`:

```go
package strs

import (
	"fmt"
	"strings"
)

type Builder struct {
	builder strings.Builder
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) Append(value any) *Builder {
	_, _ = fmt.Fprint(&b.builder, value)
	return b
}

func (b *Builder) AppendString(value string) *Builder {
	b.builder.WriteString(value)
	return b
}

func (b *Builder) AppendByte(value byte) *Builder {
	_ = b.builder.WriteByte(value)
	return b
}

func (b *Builder) String() string {
	return b.builder.String()
}

func (b *Builder) Len() int {
	return b.builder.Len()
}
```

Create `strs/joiner.go`:

```go
package strs

import "strconv"

func JoinSlice[T any](items []T, adapt func(T) string, prefix, delimiter, suffix string) string {
	b := NewBuilder().AppendString(prefix)
	for i, item := range items {
		if i > 0 {
			b.AppendString(delimiter)
		}
		b.AppendString(adapt(item))
	}
	b.AppendString(suffix)
	return b.String()
}

type Joiner[T any] struct {
	prefix    string
	delimiter string
	suffix    string
	adapt     func(T) string
}

func NewJoiner[T any](delimiter string, adapt func(T) string) *Joiner[T] {
	return &Joiner[T]{delimiter: delimiter, adapt: adapt}
}

func (j *Joiner[T]) WithPrefix(prefix string) *Joiner[T] {
	j.prefix = prefix
	return j
}

func (j *Joiner[T]) WithSuffix(suffix string) *Joiner[T] {
	j.suffix = suffix
	return j
}

func (j *Joiner[T]) Join(items []T) string {
	return JoinSlice(items, j.adapt, j.prefix, j.delimiter, j.suffix)
}

func NewStringJoiner(delimiter string) *Joiner[string] {
	return NewJoiner(delimiter, func(value string) string { return value })
}

func NewIntJoiner(delimiter string) *Joiner[int] {
	return NewJoiner(delimiter, strconv.Itoa)
}

func NewInt32Joiner(delimiter string) *Joiner[int32] {
	return NewJoiner(delimiter, func(value int32) string { return strconv.FormatInt(int64(value), 10) })
}

func NewInt64Joiner(delimiter string) *Joiner[int64] {
	return NewJoiner(delimiter, func(value int64) string { return strconv.FormatInt(value, 10) })
}

func NewUintJoiner(delimiter string) *Joiner[uint] {
	return NewJoiner(delimiter, func(value uint) string { return strconv.FormatUint(uint64(value), 10) })
}

func NewUint64Joiner(delimiter string) *Joiner[uint64] {
	return NewJoiner(delimiter, func(value uint64) string { return strconv.FormatUint(value, 10) })
}
```

Create `types/bitflag.go`:

```go
package types

type unsignedFlag interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type BitFlag[T unsignedFlag] T

func (f *BitFlag[T]) Set(flag T) {
	*f = BitFlag[T](T(*f) | flag)
}

func (f *BitFlag[T]) Clear(flag T) {
	*f = BitFlag[T](T(*f) &^ flag)
}

func (f *BitFlag[T]) Toggle(flag T) {
	*f = BitFlag[T](T(*f) ^ flag)
}

func (f BitFlag[T]) Has(flag T) bool {
	return T(f)&flag == flag
}

func (f BitFlag[T]) Value() T {
	return T(f)
}
```

- [ ] **Step 4: Run tests and verify GREEN**

Run:

```bash
go test ./strs ./types
```

Expected: PASS.

- [ ] **Step 5: Commit strs and types**

```bash
git add strs types
git commit -m "feat: add string and bitflag helpers"
```

---

### Task 5: Add concurrent task helpers

**Files:**
- Create: `concurrent/task_group.go`
- Create: `concurrent/batch.go`
- Create: `concurrent/task_group_test.go`
- Create: `concurrent/batch_test.go`

- [ ] **Step 1: Write failing concurrent tests**

Create `concurrent/task_group_test.go`:

```go
package concurrent

import (
	"context"
	"errors"
	"testing"
)

func TestTaskGroupWaitCollectsResults(t *testing.T) {
	tg := NewTaskGroup[int](context.Background())
	tg.Go(func(ctx context.Context) (int, error) { return 1, nil })
	tg.Go(func(ctx context.Context) (int, error) { return 2, nil })
	got, err := tg.Wait()
	if err != nil {
		t.Fatalf("Wait error = %v", err)
	}
	if len(got) != 2 || got[0] != 1 || got[1] != 2 {
		t.Fatalf("Wait results = %#v", got)
	}
}

func TestTaskGroupReturnsFirstErrorAndCancelsContext(t *testing.T) {
	boom := errors.New("boom")
	tg := NewTaskGroup[int](context.Background())
	tg.Go(func(ctx context.Context) (int, error) { return 0, boom })
	tg.Go(func(ctx context.Context) (int, error) {
		<-ctx.Done()
		return 0, ctx.Err()
	})
	_, err := tg.Wait()
	if !errors.Is(err, boom) {
		t.Fatalf("Wait error = %v, want boom", err)
	}
}
```

Create `concurrent/batch_test.go`:

```go
package concurrent

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

func TestAwaitTaskPreservesInputOrder(t *testing.T) {
	got, err := AwaitTask(context.Background(), func(ctx context.Context, in int) (int, error) {
		return in * 10, nil
	}, []int{3, 1, 2})
	if err != nil {
		t.Fatalf("AwaitTask error = %v", err)
	}
	if !reflect.DeepEqual(got, []int{30, 10, 20}) {
		t.Fatalf("AwaitTask = %#v", got)
	}
}

func TestAwaitBatchTaskAndDefaultBatchSize(t *testing.T) {
	got, err := AwaitBatchTask(context.Background(), func(ctx context.Context, in []int) ([]string, error) {
		out := make([]string, 0, len(in))
		for range in {
			out = append(out, "x")
		}
		return out, nil
	}, []int{1, 2, 3}, 0)
	if err != nil {
		t.Fatalf("AwaitBatchTask error = %v", err)
	}
	if !reflect.DeepEqual(got, []string{"x", "x", "x"}) {
		t.Fatalf("AwaitBatchTask = %#v", got)
	}
}

func TestAwaitBatchReturnMapTask(t *testing.T) {
	got, err := AwaitBatchReturnMapTask(context.Background(), func(ctx context.Context, in []int) (map[int]string, error) {
		out := make(map[int]string, len(in))
		for _, v := range in {
			out[v] = "ok"
		}
		return out, nil
	}, []int{1, 2, 3}, 2)
	if err != nil {
		t.Fatalf("AwaitBatchReturnMapTask error = %v", err)
	}
	if !reflect.DeepEqual(got, map[int]string{1: "ok", 2: "ok", 3: "ok"}) {
		t.Fatalf("AwaitBatchReturnMapTask = %#v", got)
	}
}

func TestAwaitTaskReturnsError(t *testing.T) {
	boom := errors.New("boom")
	_, err := AwaitTask(context.Background(), func(ctx context.Context, in int) (int, error) {
		if in == 2 {
			return 0, boom
		}
		return in, nil
	}, []int{1, 2, 3})
	if !errors.Is(err, boom) {
		t.Fatalf("AwaitTask error = %v, want boom", err)
	}
}
```

- [ ] **Step 2: Run concurrent tests and verify RED**

Run:

```bash
go test ./concurrent
```

Expected: FAIL because `concurrent` package does not exist.

- [ ] **Step 3: Implement task group and batch helpers**

Create `concurrent/task_group.go`:

```go
package concurrent

import (
	"context"
	"sync"
)

type TaskGroup[OUT any] struct {
	ctx    context.Context
	cancel context.CancelFunc
	mu     sync.Mutex
	tasks  []func(context.Context) (OUT, error)
}

func NewTaskGroup[OUT any](ctx context.Context) *TaskGroup[OUT] {
	if ctx == nil {
		ctx = context.Background()
	}
	child, cancel := context.WithCancel(ctx)
	return &TaskGroup[OUT]{ctx: child, cancel: cancel}
}

func (g *TaskGroup[OUT]) Go(fn func(context.Context) (OUT, error)) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.tasks = append(g.tasks, fn)
}

func (g *TaskGroup[OUT]) Wait() ([]OUT, error) {
	g.mu.Lock()
	tasks := append([]func(context.Context) (OUT, error)(nil), g.tasks...)
	g.mu.Unlock()

	defer g.cancel()
	results := make([]OUT, len(tasks))
	var wg sync.WaitGroup
	var errMu sync.Mutex
	var firstErr error

	for i, task := range tasks {
		i, task := i, task
		wg.Add(1)
		go func() {
			defer wg.Done()
			result, err := task(g.ctx)
			if err != nil {
				errMu.Lock()
				if firstErr == nil {
					firstErr = err
					g.cancel()
				}
				errMu.Unlock()
				return
			}
			results[i] = result
		}()
	}
	wg.Wait()
	if firstErr != nil {
		return nil, firstErr
	}
	return results, nil
}
```

Create `concurrent/batch.go`:

```go
package concurrent

import "context"

func AwaitTask[IN any, OUT any](ctx context.Context, fn func(context.Context, IN) (OUT, error), inputs []IN) ([]OUT, error) {
	tg := NewTaskGroup[OUT](ctx)
	for _, input := range inputs {
		input := input
		tg.Go(func(ctx context.Context) (OUT, error) {
			return fn(ctx, input)
		})
	}
	return tg.Wait()
}

func AwaitBatchTask[IN any, OUT any](ctx context.Context, fn func(context.Context, []IN) ([]OUT, error), inputs []IN, batchSize int) ([]OUT, error) {
	if batchSize <= 0 {
		batchSize = 1
	}
	batches := chunk(inputs, batchSize)
	tg := NewTaskGroup[[]OUT](ctx)
	for _, batch := range batches {
		batch := batch
		tg.Go(func(ctx context.Context) ([]OUT, error) {
			return fn(ctx, batch)
		})
	}
	parts, err := tg.Wait()
	if err != nil {
		return nil, err
	}
	out := make([]OUT, 0, len(inputs))
	for _, part := range parts {
		out = append(out, part...)
	}
	return out, nil
}

func AwaitBatchReturnMapTask[IN comparable, OUT any](ctx context.Context, fn func(context.Context, []IN) (map[IN]OUT, error), inputs []IN, batchSize int) (map[IN]OUT, error) {
	if batchSize <= 0 {
		batchSize = 1
	}
	batches := chunk(inputs, batchSize)
	tg := NewTaskGroup[map[IN]OUT](ctx)
	for _, batch := range batches {
		batch := batch
		tg.Go(func(ctx context.Context) (map[IN]OUT, error) {
			return fn(ctx, batch)
		})
	}
	parts, err := tg.Wait()
	if err != nil {
		return nil, err
	}
	out := make(map[IN]OUT, len(inputs))
	for _, part := range parts {
		for key, value := range part {
			out[key] = value
		}
	}
	return out, nil
}

func chunk[T any](items []T, size int) [][]T {
	if items == nil {
		return nil
	}
	out := make([][]T, 0, (len(items)+size-1)/size)
	for start := 0; start < len(items); start += size {
		end := start + size
		if end > len(items) {
			end = len(items)
		}
		out = append(out, items[start:end])
	}
	return out
}
```

- [ ] **Step 4: Run concurrent tests and verify GREEN**

Run:

```bash
go test ./concurrent
```

Expected: PASS.

- [ ] **Step 5: Commit concurrent package**

```bash
git add concurrent
git commit -m "feat: add concurrent task helpers"
```

---

### Task 6: Update README and verify whole repository

**Files:**
- Modify: `README.md`

- [ ] **Step 1: Replace README content**

Write `README.md`:

```markdown
# candy

Go utility packages for common application code.

## Packages

- `arrays`: generic slice filtering, mapping, reducing, chunking, grouping, and deduplication.
- `optional`: explicit-present Optional values.
- `cache`: concurrent-safe LRU cache.
- `concurrent`: context-aware task groups and batch helpers.
- `strs`: chainable string builder and generic joiners.
- `types`: lightweight type helpers such as bit flags.

The repository also contains earlier experimental Java-style collection and stream packages. New utility packages are designed to be independent and practical for normal Go code.

## Requirements

Go 1.18 or newer.

## Test

```bash
go test ./...
```
```

- [ ] **Step 2: Run gofmt**

Run:

```bash
gofmt -w arrays optional cache strs types concurrent
```

Expected: command exits 0.

- [ ] **Step 3: Run full test suite**

Run:

```bash
go test ./...
```

Expected: PASS. If existing unfinished packages fail independently of the new work, record the exact failures, fix only compile breakages introduced by this plan, and do not expand scope into stream/collection feature completion.

- [ ] **Step 4: Commit README and any verification fixes**

```bash
git add README.md
git commit -m "docs: describe utility packages"
```

- [ ] **Step 5: Final status check**

Run:

```bash
git status --short
go test ./...
```

Expected: clean working tree and passing tests.

---

## Self-Review Notes

- The plan covers every approved spec package: arrays, optional, cache, concurrent, strs, and types.
- It keeps Go 1.18 compatibility: no `slices`, `maps`, method type parameters, or newer language features.
- It avoids destructive rewrites except the approved optional simplification.
- Tests are written before implementation in every feature task.
