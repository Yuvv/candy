package stream

import (
	"math"
	"reflect"
	"testing"

	"github.com/yuvv/candy/function"
)

type intComparator struct{}

func (intComparator) Compare(a, b int) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

func (intComparator) Equals(obj any) bool {
	_, ok := obj.(intComparator)
	return ok
}

func TestOfCopiesInputAndToArrayReturnsCopy(t *testing.T) {
	values := []int{3, 1, 2}
	stream := Of(values...)
	values[0] = 99

	got := stream.ToArray()
	if want := []int{3, 1, 2}; !reflect.DeepEqual(got, want) {
		t.Fatalf("ToArray() = %v, want %v", got, want)
	}

	got[1] = 88
	if want := []int{3, 1, 2}; !reflect.DeepEqual(stream.ToArray(), want) {
		t.Fatalf("ToArray() after returned slice mutation = %v, want %v", stream.ToArray(), want)
	}
}

func TestEmptyTerminalOperations(t *testing.T) {
	stream := Empty[int]()
	if got := stream.Count(); got != 0 {
		t.Fatalf("Count() = %d, want 0", got)
	}
	if got := stream.ToArray(); len(got) != 0 {
		t.Fatalf("ToArray() = %v, want empty", got)
	}
	if got := stream.FindFirst(); !got.IsEmpty() {
		t.Fatalf("FindFirst() = %v, want empty", got)
	}
	if got := stream.Reduce(func(a, b int) int { return a + b }); got != 0 {
		t.Fatalf("Reduce() = %d, want zero", got)
	}
}

func TestConcatStreamsAndNilOperands(t *testing.T) {
	tests := []struct {
		name string
		a, b Stream[int]
		want []int
	}{
		{name: "two streams", a: Of(1, 2), b: Of(3, 4), want: []int{1, 2, 3, 4}},
		{name: "nil first", b: Of(3, 4), want: []int{3, 4}},
		{name: "nil second", a: Of(1, 2), want: []int{1, 2}},
		{name: "both nil", want: []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Concat(tt.a, tt.b).ToArray(); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("Concat().ToArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForEachPreservesOrder(t *testing.T) {
	var got []int
	Of(3, 1, 2).ForEach(func(value int) { got = append(got, value) })
	if want := []int{3, 1, 2}; !reflect.DeepEqual(got, want) {
		t.Fatalf("ForEach order = %v, want %v", got, want)
	}
}

func TestForEachOrderedPreservesOrder(t *testing.T) {
	var got []int
	Of(3, 1, 2).ForEachOrdered(func(value int) { got = append(got, value) })
	if want := []int{3, 1, 2}; !reflect.DeepEqual(got, want) {
		t.Fatalf("ForEachOrdered order = %v, want %v", got, want)
	}
}

func TestReduceAndReduceWithIdentity(t *testing.T) {
	add := function.BinaryOperator[int](func(a, b int) int { return a + b })
	if got := Of(1, 2, 3, 4).Reduce(add); got != 10 {
		t.Fatalf("Reduce() = %d, want 10", got)
	}
	if got := Of(1, 2, 3, 4).ReduceWithIdentity(10, add); got != 20 {
		t.Fatalf("ReduceWithIdentity() = %d, want 20", got)
	}
	if got := Empty[int]().ReduceWithIdentity(10, add); got != 10 {
		t.Fatalf("empty ReduceWithIdentity() = %d, want identity 10", got)
	}
}

func TestCount(t *testing.T) {
	if got := Of("a", "b", "c").Count(); got != 3 {
		t.Fatalf("Count() = %d, want 3", got)
	}
}

func TestAnyMatch(t *testing.T) {
	isEven := function.Predicate[int](func(value int) bool { return value%2 == 0 })

	if !Of(1, 2, 3).AnyMatch(isEven) {
		t.Fatal("AnyMatch() = false, want true for matching stream")
	}
	if Of(1, 3, 5).AnyMatch(isEven) {
		t.Fatal("AnyMatch() = true, want false for non-matching stream")
	}
	if Empty[int]().AnyMatch(isEven) {
		t.Fatal("empty AnyMatch() = true, want false")
	}

	calls := 0
	if !Of(1, 2, 4).AnyMatch(func(value int) bool {
		calls++
		return value%2 == 0
	}) {
		t.Fatal("AnyMatch() = false, want true")
	}
	if calls != 2 {
		t.Fatalf("AnyMatch() predicate calls = %d, want 2", calls)
	}
}

func TestAllMatch(t *testing.T) {
	isEven := function.Predicate[int](func(value int) bool { return value%2 == 0 })

	if !Of(2, 4, 6).AllMatch(isEven) {
		t.Fatal("AllMatch() = false, want true for matching stream")
	}
	if Of(2, 3, 4).AllMatch(isEven) {
		t.Fatal("AllMatch() = true, want false for non-matching stream")
	}
	if !Empty[int]().AllMatch(isEven) {
		t.Fatal("empty AllMatch() = false, want true")
	}

	calls := 0
	if Of(2, 3, 4).AllMatch(func(value int) bool {
		calls++
		return value%2 == 0
	}) {
		t.Fatal("AllMatch() = true, want false")
	}
	if calls != 2 {
		t.Fatalf("AllMatch() predicate calls = %d, want 2", calls)
	}
}

func TestNoneMatch(t *testing.T) {
	isEven := function.Predicate[int](func(value int) bool { return value%2 == 0 })

	if !Of(1, 3, 5).NoneMatch(isEven) {
		t.Fatal("NoneMatch() = false, want true for non-matching stream")
	}
	if Of(1, 2, 3).NoneMatch(isEven) {
		t.Fatal("NoneMatch() = true, want false for matching stream")
	}
	if !Empty[int]().NoneMatch(isEven) {
		t.Fatal("empty NoneMatch() = false, want true")
	}

	calls := 0
	if Of(1, 2, 4).NoneMatch(func(value int) bool {
		calls++
		return value%2 == 0
	}) {
		t.Fatal("NoneMatch() = true, want false")
	}
	if calls != 2 {
		t.Fatalf("NoneMatch() predicate calls = %d, want 2", calls)
	}
}

func TestFindFirstAndFindAny(t *testing.T) {
	stream := Of(7, 8, 9)
	if got := stream.FindFirst(); !got.IsPresent() || got.Get() != 7 {
		t.Fatalf("FindFirst() = %v, want 7", got)
	}
	if got := stream.FindAny(); !got.IsPresent() || got.Get() != 7 {
		t.Fatalf("FindAny() = %v, want first value 7", got)
	}
	if !Empty[int]().FindAny().IsEmpty() {
		t.Fatal("empty FindAny() should be empty")
	}
}

func TestFilterPreservesOrderAndSource(t *testing.T) {
	source := Of(3, 2, 4, 1)
	filtered := source.Filter(func(value int) bool { return value%2 == 0 })

	if got, want := filtered.ToArray(), []int{2, 4}; !reflect.DeepEqual(got, want) {
		t.Fatalf("Filter().ToArray() = %v, want %v", got, want)
	}
	if got, want := source.ToArray(), []int{3, 2, 4, 1}; !reflect.DeepEqual(got, want) {
		t.Fatalf("source after Filter() = %v, want %v", got, want)
	}
}

func TestDistinctPreservesFirstOccurrence(t *testing.T) {
	if got, want := Of(3, 1, 3, 2, 1).Distinct().ToArray(), []int{3, 1, 2}; !reflect.DeepEqual(got, want) {
		t.Fatalf("Distinct().ToArray() = %v, want %v", got, want)
	}
}

func TestDistinctUsesDeepEqualForNonComparableValues(t *testing.T) {
	type record struct {
		name string
		tags []string
	}
	values := []record{
		{name: "first", tags: []string{"a", "b"}},
		{name: "second", tags: []string{"x"}},
		{name: "first", tags: []string{"a", "b"}},
	}
	want := values[:2]
	if got := Of(values...).Distinct().ToArray(); !reflect.DeepEqual(got, want) {
		t.Fatalf("Distinct().ToArray() = %v, want %v", got, want)
	}

	slices := [][]int{{1, 2}, {3}, {1, 2}}
	if got, want := Of(slices...).Distinct().ToArray(), slices[:2]; !reflect.DeepEqual(got, want) {
		t.Fatalf("slice Distinct().ToArray() = %v, want %v", got, want)
	}
}

func TestPeekCallsConsumerInOrderAndReturnsIndependentStream(t *testing.T) {
	source := Of(3, 1, 2)
	var calls []int
	peeked := source.Peek(func(value int) { calls = append(calls, value) })

	if want := []int{3, 1, 2}; !reflect.DeepEqual(calls, want) {
		t.Fatalf("Peek() call order = %v, want %v", calls, want)
	}
	if got, want := peeked.ToArray(), []int{3, 1, 2}; !reflect.DeepEqual(got, want) {
		t.Fatalf("Peek().ToArray() = %v, want %v", got, want)
	}

	peekedValue := peeked.(sequentialStream[int])
	peekedValue.values[0] = 99
	if got, want := source.ToArray(), []int{3, 1, 2}; !reflect.DeepEqual(got, want) {
		t.Fatalf("source after mutating Peek result = %v, want %v", got, want)
	}
}

func TestSkip(t *testing.T) {
	tests := []struct {
		name string
		n    int64
		want []int
	}{
		{name: "negative", n: -1, want: []int{1, 2, 3}},
		{name: "zero", n: 0, want: []int{1, 2, 3}},
		{name: "normal", n: 2, want: []int{3}},
		{name: "equal count", n: 3, want: []int{}},
		{name: "too large", n: 10, want: []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Of(1, 2, 3).Skip(tt.n).ToArray(); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("Skip(%d).ToArray() = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func TestLimit(t *testing.T) {
	assertPanicMessage(t, "stream: limit size must be non-negative", func() {
		Of(1, 2, 3).Limit(-1)
	})

	tests := []struct {
		name string
		n    int64
		want []int
	}{
		{name: "zero", n: 0, want: []int{}},
		{name: "normal", n: 2, want: []int{1, 2}},
		{name: "equal count", n: 3, want: []int{1, 2, 3}},
		{name: "too large", n: 10, want: []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Of(1, 2, 3).Limit(tt.n).ToArray(); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("Limit(%d).ToArray() = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func TestSortedOrderedTypesAndSourceUnchanged(t *testing.T) {
	ints := Of(3, 1, 2)
	if got, want := ints.Sorted().ToArray(), []int{1, 2, 3}; !reflect.DeepEqual(got, want) {
		t.Fatalf("int Sorted().ToArray() = %v, want %v", got, want)
	}
	if got, want := ints.ToArray(), []int{3, 1, 2}; !reflect.DeepEqual(got, want) {
		t.Fatalf("int source after Sorted() = %v, want %v", got, want)
	}

	if got, want := Of("beta", "alpha", "gamma").Sorted().ToArray(), []string{"alpha", "beta", "gamma"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("string Sorted().ToArray() = %v, want %v", got, want)
	}
	if got, want := Of(3.5, -1.0, 2.25).Sorted().ToArray(), []float64{-1.0, 2.25, 3.5}; !reflect.DeepEqual(got, want) {
		t.Fatalf("float Sorted().ToArray() = %v, want %v", got, want)
	}
}

func TestSortedFloat64PlacesNaNsLast(t *testing.T) {
	firstNaN := math.Float64frombits(0x7ff8000000000001)
	secondNaN := math.Float64frombits(0x7ff8000000000002)
	got := Of(firstNaN, -1.0, secondNaN, 0.0, 2.0).Sorted().ToArray()
	if want := []float64{-1, 0, 2}; !reflect.DeepEqual(got[:3], want) {
		t.Fatalf("Sorted().ToArray() numeric values = %v, want %v", got[:3], want)
	}
	if !math.IsNaN(got[3]) || !math.IsNaN(got[4]) {
		t.Fatalf("Sorted().ToArray() trailing values = %v, want NaNs", got[3:])
	}
	if gotBits, wantBits := []uint64{math.Float64bits(got[3]), math.Float64bits(got[4])}, []uint64{math.Float64bits(firstNaN), math.Float64bits(secondNaN)}; !reflect.DeepEqual(gotBits, wantBits) {
		t.Fatalf("Sorted().ToArray() NaN order = %x, want stable order %x", gotBits, wantBits)
	}
}

func TestSortedFloat32PlacesNaNsLast(t *testing.T) {
	nan := float32(math.NaN())
	got := Of(nan, float32(-1), float32(0), float32(2)).Sorted().ToArray()
	if want := []float32{-1, 0, 2}; !reflect.DeepEqual(got[:3], want) {
		t.Fatalf("Sorted().ToArray() numeric values = %v, want %v", got[:3], want)
	}
	if !math.IsNaN(float64(got[3])) {
		t.Fatalf("Sorted().ToArray() last value = %v, want NaN", got[3])
	}
}

func TestSortedUnsupportedTypePanics(t *testing.T) {
	assertPanicMessage(t, "stream: Sorted requires an ordered element type", func() {
		Of(struct{ value int }{value: 1}).Sorted()
	})
}

func TestSortedByMethodPanicsWithPackageFunctionGuidance(t *testing.T) {
	assertPanicMessage(t, "stream: SortedBy requires a comparator; use stream.SortedBy", func() {
		Of(3, 1, 2).SortedBy()
	})
}

func TestMinAndMax(t *testing.T) {
	stream := Of(4, 2, 9, 1, 7)
	comparator := intComparator{}
	if got := stream.Min(comparator); !got.IsPresent() || got.Get() != 1 {
		t.Fatalf("Min() = %v, want 1", got)
	}
	if got := stream.Max(comparator); !got.IsPresent() || got.Get() != 9 {
		t.Fatalf("Max() = %v, want 9", got)
	}
	if !Empty[int]().Min(comparator).IsEmpty() {
		t.Fatal("empty Min() should be empty")
	}
	if !Empty[int]().Max(comparator).IsEmpty() {
		t.Fatal("empty Max() should be empty")
	}
}
