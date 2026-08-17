package stream

import (
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
