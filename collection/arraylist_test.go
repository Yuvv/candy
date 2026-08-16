package collection

import (
	"fmt"
	"strings"
	"testing"
)

func TestCollection(t *testing.T) {
	list := NewArrayList[int]()
	if list == nil {
		t.Fatal("list cannot be nil")
	}
	// check empty
	if !list.IsEmpty() {
		t.Fatalf("list should be empty after initialized")
	}
	// check Add+Size
	list.Add(1)
	list.Add(2)
	if list.Size() != 2 {
		t.Fatalf("size must be 2 after add 2 element")
	}
	// check contains
	if list.Contains(4) {
		t.Fatalf("list doesn't contains `4` but return true")
	}
	if !list.Contains(2) {
		t.Fatalf("list contains `2` but return false")
	}
	// check toArray
	la := list.ToArray()
	if la == nil || len(la) != 2 || la[0] != 1 || la[1] != 2 {
		t.Fatalf("ToArray returns wrong result, expected:[1,2], but got:%+v", la)
	}
	// check Remove
	list.Remove(2)
	if list.Size() != 1 || list.Contains(2) {
		t.Fatalf("Remove failed")
	}
	// check ContainsAll
	l2 := NewArrayList[int64]()
	l2.Add(1)
	l2.Add(3)
	l2.Add(5)
	l2.Add(7)
	l2.Add(9)
	if !l2.ContainsAll(NewArrayListWithEle[int64](1, 3, 5, 7, 9)) {
		t.Fatal("ContainsAll faile")
	}
	// check AddAll
	l3 := NewArrayListWithEle[int64](2, 4, 6, 8, 10)
	modified := l2.AddAll(l3)
	if !modified {
		t.Fatal("should modified after add elements")
	}
	if !l2.ContainsAll(l3) {
		t.Fatal("should contains all elements after AddAll")
	}
	if l2.Size() != 10 {
		t.Fatal("size should be 10 after add 5 elements")
	}
	// check RemoveAll
	modified = l2.RemoveAll(NewArrayListWithEle[int64](1, 3, 5, 7, 9))
	if !modified {
		t.Fatal("should modified after remove elements")
	}
	if l2.Size() != 5 {
		t.Fatal("size should be 5 after remove 5 elements")
	}
	if !l2.ContainsAll(l3) {
		t.Fatalf("list should contains all even number after remove odd numbers.\n got l2:%+v", l2)
	}
	// check RemoveIf
	modified = l2.RemoveIf(func(it int64) bool {
		return it%2 == 0
	})
	if !modified {
		t.Fatal("should modified after RemoveIf elements")
	}
	if l2.Size() != 0 {
		t.Fatal("size should be 0 after remove all odd number")
	}
	// check RetailAll
	l2.AddAll(l3)
	modified = l2.RetainAll(NewArrayListWithEle[int64](1, 4, 9, 8, -1))
	if !modified {
		t.Fatal("should modified after RetainAll")
	}
	if l2.Size() != 2 {
		t.Fatalf("size should be 2 (got %d) after RetainAll has 2 intersection", l2.Size())
	}
	modified = l2.RetainAll(NewArrayListWithEle[int64](4, 8))
	if modified {
		t.Fatal("should not modified after RetailAll has no intersection")
	}
	if !l2.ContainsAll(NewArrayListWithEle[int64](4, 8)) {
		t.Fatalf("list should be [2,4,6,10] but got:%+v", l2.slice)
	}
	// check Clear
	l2.Clear()
	if l2.Size() != 0 {
		t.Fatal("size must be 0 after Clear")
	}
}

func Test_ArrayList_AddAt(t *testing.T) {
	lst := NewArrayListWithEle[int](1, 2, 3, 4, 5)
	originSize := lst.Size()
	lst.AddAt(0, 0)
	if lst.Size() != originSize+1 {
		t.Fatalf("size should be %d after add 1 element", originSize+1)
	}
	for i := 0; i < lst.Size(); i++ {
		if lst.Get(i) != i {
			t.Fatalf("element at `%d` should be `%d`", i, i)
		}
	}

	originSize = lst.Size()
	lst.AddAt(lst.Size(), 6)
	if lst.Size() != originSize+1 {
		t.Fatalf("size should be %d after add 1 element", originSize+1)
	}
	for i := 0; i < lst.Size(); i++ {
		if lst.Get(i) != i {
			t.Fatalf("element at `%d` should be `%d`", i, i)
		}
	}

	originSize = lst.Size()
	lst.AddAt(2, 2)
	if lst.Size() != originSize+1 {
		t.Fatalf("size should be %d after add 1 element", originSize+1)
	}
	for i := 0; i < lst.Size(); i++ {
		if i <= 2 && lst.Get(i) != i {
			t.Fatalf("element at `%d` should be `%d`", i, i)
		} else if i > 2 && lst.Get(i) != i-1 {
			t.Fatalf("element at `%d` should be `%d`", i, i-1)
		}
	}

}

func assertPanicsWithIndex(t *testing.T, wantIdx int, action func()) {
	t.Helper()
	defer func() {
		recovered := recover()
		if recovered == nil {
			t.Fatal("expected panic")
		}
		message := fmt.Sprint(recovered)
		if !strings.Contains(message, "index") || !strings.Contains(message, fmt.Sprint(wantIdx)) {
			t.Fatalf("panic %q should clearly identify invalid index %d", message, wantIdx)
		}
	}()
	action()
}

func pointerList(values ...*int) *_ArrayList[*int] {
	return NewSpecArrayListWithEle(func(a *int, o any) bool {
		other, ok := o.(*int)
		return ok && a == other
	}, values...)
}

func assertDiscardedPointersCleared(t *testing.T, lst *_ArrayList[*int]) {
	t.Helper()
	full := lst.slice[:cap(lst.slice)]
	for i := len(lst.slice); i < len(full); i++ {
		if full[i] != nil {
			t.Fatalf("discarded slot %d retains pointer %p", i, full[i])
		}
	}
}

func Test_ArrayList_AddAtEmptyAndAppend(t *testing.T) {
	lst := NewArrayList[int]()
	lst.AddAt(0, 1)
	lst.AddAt(lst.Size(), 2)

	if got := lst.ToArray(); len(got) != 2 || got[0] != 1 || got[1] != 2 {
		t.Fatalf("AddAt results = %v, want [1 2]", got)
	}
}

func Test_ArrayList_AddAtRejectsInvalidIndex(t *testing.T) {
	for _, idx := range []int{-1, 2} {
		t.Run(fmt.Sprint(idx), func(t *testing.T) {
			lst := NewArrayListWithEle(1)
			assertPanicsWithIndex(t, idx, func() { lst.AddAt(idx, 2) })
		})
	}
}

func Test_ArrayList_RemoveAtRejectsInvalidIndex(t *testing.T) {
	for _, idx := range []int{-1, 1} {
		t.Run(fmt.Sprint(idx), func(t *testing.T) {
			lst := NewArrayListWithEle(1)
			assertPanicsWithIndex(t, idx, func() { lst.RemoveAt(idx) })
		})
	}
}

func Test_ArrayList_RemovalsClearDiscardedPointerSlots(t *testing.T) {
	t.Run("Remove", func(t *testing.T) {
		a, b, c := 1, 2, 3
		lst := pointerList(&a, &b, &c)
		lst.Remove(&b)
		assertDiscardedPointersCleared(t, lst)
	})

	t.Run("RemoveAt", func(t *testing.T) {
		a, b, c := 1, 2, 3
		lst := pointerList(&a, &b, &c)
		lst.RemoveAt(0)
		assertDiscardedPointersCleared(t, lst)
	})

	t.Run("RemoveIf", func(t *testing.T) {
		a, b, c, d := 1, 2, 3, 4
		lst := pointerList(&a, &b, &c, &d)
		lst.RemoveIf(func(value *int) bool { return *value%2 == 0 })
		assertDiscardedPointersCleared(t, lst)
	})

	t.Run("Clear", func(t *testing.T) {
		a, b, c := 1, 2, 3
		lst := pointerList(&a, &b, &c)
		lst.Clear()
		assertDiscardedPointersCleared(t, lst)
	})
}

func Test_ArrayList_AddAllAtValidatesIndex(t *testing.T) {
	for _, idx := range []int{-1, 2} {
		t.Run(fmt.Sprint(idx), func(t *testing.T) {
			lst := NewArrayListWithEle(1)
			assertPanicsWithIndex(t, idx, func() { lst.AddAllAt(idx, NewArrayList[int]()) })
		})
	}
}

func Test_ArrayList_AddAllAtEmptyCollectionDoesNotModify(t *testing.T) {
	lst := NewArrayListWithEle(1)
	itr := lst.Iterator()

	lst.AddAllAt(1, NewArrayList[int]())

	if got := itr.Next(); got != 1 {
		t.Fatalf("iterator returned %d, want 1", got)
	}
}

func Test_ArrayList_ClearEmptyDoesNotModify(t *testing.T) {
	lst := NewArrayList[int]()
	before := lst.modCount

	lst.Clear()

	if lst.modCount != before {
		t.Fatalf("Clear modCount = %d, want %d", lst.modCount, before)
	}
}
