package stream

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestMapPreservesOrderAndChangesType(t *testing.T) {
	got := Map(Of(3, 1, 2), strconv.Itoa).ToArray()
	want := []string{"3", "1", "2"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Map().ToArray() = %v, want %v", got, want)
	}
}

func TestMapNilStreamIsEmpty(t *testing.T) {
	var source Stream[int]
	if got := Map(source, strconv.Itoa).ToArray(); len(got) != 0 {
		t.Fatalf("Map(nil).ToArray() = %v, want empty", got)
	}
}

func TestFlatMapExpandsInOrderAndSkipsNilStream(t *testing.T) {
	got := FlatMap(Of(1, 2, 3), func(value int) Stream[string] {
		if value == 2 {
			return nil
		}
		return Of(strconv.Itoa(value), strconv.Itoa(value*10))
	}).ToArray()
	want := []string{"1", "10", "3", "30"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("FlatMap().ToArray() = %v, want %v", got, want)
	}
}

func TestFlatMapNilStreamIsEmpty(t *testing.T) {
	var source Stream[int]
	if got := FlatMap(source, func(value int) Stream[string] { return Of(strconv.Itoa(value)) }).ToArray(); len(got) != 0 {
		t.Fatalf("FlatMap(nil).ToArray() = %v, want empty", got)
	}
}

type aliasingStream[T any] struct {
	Stream[T]
	values []T
}

func (s aliasingStream[T]) ToArray() []T {
	return s.values
}

func TestSortedBySortsCustomStructWithoutChangingSource(t *testing.T) {
	type person struct {
		name string
		age  int
	}

	sourceValues := []person{{name: "Grace", age: 37}, {name: "Ada", age: 31}, {name: "Linus", age: 34}}
	source := Of(sourceValues...)
	got := SortedBy(source, func(a, b person) bool { return a.age < b.age }).ToArray()
	want := []person{{name: "Ada", age: 31}, {name: "Linus", age: 34}, {name: "Grace", age: 37}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("SortedBy().ToArray() = %v, want %v", got, want)
	}
	if got := source.ToArray(); !reflect.DeepEqual(got, sourceValues) {
		t.Fatalf("source after SortedBy() = %v, want %v", got, sourceValues)
	}
}

func TestSortedByPreservesEqualKeyOrder(t *testing.T) {
	type item struct {
		group int
		name  string
	}

	values := make([]item, 0, 40)
	want := make([]item, 0, 40)
	for i := 0; i < 20; i++ {
		values = append(values,
			item{group: 2, name: fmt.Sprintf("second-%02d", i)},
			item{group: 1, name: fmt.Sprintf("first-%02d", i)},
		)
	}
	for i := 0; i < 20; i++ {
		want = append(want, item{group: 1, name: fmt.Sprintf("first-%02d", i)})
	}
	for i := 0; i < 20; i++ {
		want = append(want, item{group: 2, name: fmt.Sprintf("second-%02d", i)})
	}
	got := SortedBy(Of(values...), func(a, b item) bool { return a.group < b.group }).ToArray()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("SortedBy().ToArray() = %v, want stable order %v", got, want)
	}
}

func TestSortedByDoesNotMutateAliasingSource(t *testing.T) {
	values := []int{3, 1, 2}
	source := aliasingStream[int]{Stream: Empty[int](), values: values}
	if got, want := SortedBy[int](source, func(a, b int) bool { return a < b }).ToArray(), []int{1, 2, 3}; !reflect.DeepEqual(got, want) {
		t.Fatalf("SortedBy().ToArray() = %v, want %v", got, want)
	}
	if want := []int{3, 1, 2}; !reflect.DeepEqual(values, want) {
		t.Fatalf("source backing values after SortedBy() = %v, want %v", values, want)
	}
}

func TestSortedByNilStreamIsEmpty(t *testing.T) {
	var source Stream[int]
	if got := SortedBy(source, func(a, b int) bool { return a < b }).ToArray(); len(got) != 0 {
		t.Fatalf("SortedBy(nil).ToArray() = %v, want empty", got)
	}
}

func TestSortedByNilLessPanics(t *testing.T) {
	assertPanicMessage(t, "stream: SortedBy requires a less function", func() {
		SortedBy[int](Of(1, 2, 3), nil)
	})
}

func TestMapChainsWithStreamMethods(t *testing.T) {
	got := Map(Of(1, 2, 3).Filter(func(value int) bool { return value%2 == 1 }), strconv.Itoa).ToArray()
	want := []string{"1", "3"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("chained Map().ToArray() = %v, want %v", got, want)
	}
}
