package collection

import "testing"

func TestHashSetIsEmpty(t *testing.T) {
	set := NewHashSet[int]()
	if !set.IsEmpty() {
		t.Fatal("new set is not empty")
	}

	set.Add(1)
	if set.IsEmpty() {
		t.Fatal("set with an element is empty")
	}
}

func TestHashSetIteratorVisitsSnapshotValues(t *testing.T) {
	set := NewHashSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	it := set.Iterator()
	set.Add(4)

	got := make(map[int]int)
	it.ForEachRemaining(func(value int) {
		got[value]++
	})

	if len(got) != 3 {
		t.Fatalf("iterator visited %d distinct values, want 3: %v", len(got), got)
	}
	for _, value := range []int{1, 2, 3} {
		if got[value] != 1 {
			t.Errorf("iterator visited %d %d times, want once", value, got[value])
		}
	}
	if got[4] != 0 {
		t.Error("iterator visited value added after iterator creation")
	}
}

func TestHashSetIteratorRemoveRemovesLastReturnedValue(t *testing.T) {
	set := NewHashSet[int]()
	set.Add(1)
	set.Add(2)

	it := set.Iterator()
	removed := it.Next()
	it.Remove()

	if set.Contains(removed) {
		t.Fatalf("set still contains iterator-removed value %d", removed)
	}
	if set.Size() != 1 {
		t.Fatalf("set size = %d, want 1", set.Size())
	}
}

func TestHashSetIteratorRemoveBeforeNextPanics(t *testing.T) {
	set := NewHashSet[int]()
	set.Add(1)
	it := set.Iterator()

	defer func() {
		if recover() == nil {
			t.Fatal("Remove before Next did not panic")
		}
	}()
	it.Remove()
}

func TestHashSetAddAllReportsOnlyNewElements(t *testing.T) {
	set := NewHashSet[int]()
	set.Add(1)
	set.Add(2)

	if set.AddAll(NewArrayListWithEle(1, 2, 1)) {
		t.Fatal("AddAll returned true for duplicate-only source")
	}
	if !set.AddAll(NewArrayListWithEle(2, 3, 2)) {
		t.Fatal("AddAll returned false when source contained a new value")
	}
	if !set.Contains(3) {
		t.Fatal("AddAll did not add new value")
	}
}
