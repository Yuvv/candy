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
