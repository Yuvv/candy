package lists

import "testing"

func TestCollection(t *testing.T) {
	list := NewArrayList[int]()
	if list == nil {
		t.Error("list cannot be nil")
	}
	// check empty
	if !list.IsEmpty() {
		t.Errorf("list should be empty after initialized")
	}
	// check Add+Size
	list.Add(1)
	list.Add(2)
	if list.Size() != 2 {
		t.Errorf("size must be 2 after add 2 element")
	}
	// check contains
	if list.Contains(4) {
		t.Errorf("list doesn't contains `4` but return true")
	}
	if !list.Contains(2) {
		t.Errorf("list contains `2` but return false")
	}
	// check toArray
	la := list.ToArray()
	if la == nil || len(la) != 2 || la[0] != 1 || la[1] != 2 {
		t.Errorf("ToArray returns wrong result, expected:[1,2], but got:%+v", la)
	}
	// check Remove
	list.Remove(2)
	if list.Size() != 1 || list.Contains(2) {
		t.Errorf("Remove failed")
	}
	// check ContainsAll
	l2 := NewArrayList[int64]()
	l2.Add(1)
	l2.Add(3)
	l2.Add(5)
	l2.Add(7)
	l2.Add(9)
	if !l2.ContainsAll(NewArrayListWithEle[int64](1, 3, 5, 7, 9)) {
		t.Errorf("ContainsAll faile")
	}
}

func TestList(t *testing.T) {
	//
}
