package lists

import (
	"github.com/yuvv/candy/collections"
	"github.com/yuvv/candy/function"
	"github.com/yuvv/candy/iters"
	"github.com/yuvv/candy/lang"
)

type ArrayList[E comparable] struct {
	AbstractList[E]

	slice []E
}

func (lst *ArrayList[E]) Iterator() iters.Iterator[E] {
	return NewItr(lst)
}

func (lst *ArrayList[E]) ForEach(action function.Consumer[E]) {
	for _, item := range lst.slice {
		action(item)
	}
}

func (lst *ArrayList[E]) Spliterator() iters.Spliterator[E] {
	//TODO implement me
	panic("implement me")
}

func (lst *ArrayList[E]) Size() int {
	return len(lst.slice)
}

func (lst *ArrayList[E]) IsEmpty() bool {
	return lst.Size() == 0
}

func (lst *ArrayList[E]) Contains(o E) bool {
	for _, it := range lst.slice {
		if it == o {
			return true
		}
	}
	return false
}

func (lst *ArrayList[E]) ToArray() []E {
	toReturn := make([]E, lst.Size())
	for i, item := range lst.slice {
		toReturn[i] = item
	}
	return toReturn
}

func (lst *ArrayList[E]) Add(e E) bool {
	lst.slice = append(lst.slice, e)
	return true
}

func (lst *ArrayList[E]) Remove(o E) bool {
	for i, item := range lst.slice {
		if item == o {
			for j := i + 1; j < len(lst.slice); j++ {
				lst.slice[j-1] = lst.slice[j]
			}
			lst.slice = lst.slice[:len(lst.slice)-1]
			return true
		}
	}
	return false
}

func (lst *ArrayList[E]) ContainsAll(collection collections.Collection[E]) bool {
	it := collection.Iterator()
	for it.HasNext() {
		if !lst.Contains(it.Next()) {
			return false
		}
	}
	return true
}

func (lst *ArrayList[E]) AddAll(collection collections.Collection[E]) bool {
	modified := false
	it := collection.Iterator()
	for it.HasNext() {
		modified = lst.Add(it.Next())
	}
	return modified
}

func (lst *ArrayList[E]) RemoveAll(collection collections.Collection[E]) bool {
	i := 0
	for j := 0; j < len(lst.slice); j++ {
		if collection.Contains(lst.slice[j]) {
			continue
		}
		lst.slice[i] = lst.slice[j]
		i++
	}
	if i >= len(lst.slice) {
		return false
	}
	lst.slice = lst.slice[:i]
	return true
}

func (lst *ArrayList[E]) RemoveIf(predicate function.Predicate[E]) bool {
	//TODO implement me
	panic("implement me")
}

func (lst *ArrayList[E]) RetainAll(collection collections.Collection[E]) bool {
	modified := false
	it := lst.Iterator()
	for it.HasNext() {
		if !collection.Contains(it.Next()) {
			it.Remove()
			modified = true
		}
	}
	return modified
}

func (lst *ArrayList[E]) Clear() {
	lst.slice = lst.slice[:0]
}

func (lst *ArrayList[E]) AddAllAt(idx int, collection collections.Collection[E]) {
	//TODO implement me
	panic("implement me")
}

func (lst *ArrayList[E]) ReplaceAll(operator function.UnaryOperator[E]) {
	//TODO implement me
	panic("implement me")
}

func (lst *ArrayList[E]) Sort(comparator lang.Comparator[E]) {
	//TODO implement me
	panic("implement me")
}

func (lst *ArrayList[E]) Get(idx int) E {
	//TODO implement me
	panic("implement me")
}

func (lst *ArrayList[E]) Set(idx int, ele E) E {
	//TODO implement me
	panic("implement me")
}

func (lst *ArrayList[E]) AddAt(idx int, ele E) {
	//TODO implement me
	panic("implement me")
}

func (lst *ArrayList[E]) RemoveAt(idx int) E {
	//TODO implement me
	panic("implement me")
}

func (lst *ArrayList[E]) IndexOf(o E) int {
	//TODO implement me
	panic("implement me")
}

func (lst *ArrayList[E]) LastIndexOf(o E) int {
	//TODO implement me
	panic("implement me")
}

func (lst *ArrayList[E]) ListIterator() iters.ListIterator[E] {
	//TODO implement me
	panic("implement me")
}

func (lst *ArrayList[E]) ListIteratorFrom(idx int) iters.ListIterator[E] {
	//TODO implement me
	panic("implement me")
}

func (lst *ArrayList[E]) SubList(fromIdx, toIdx int) List[E] {
	//TODO implement me
	panic("implement me")
}

func (lst *ArrayList[E]) getModCount() int {
	return lst.modCount
}

/// ---------------------------------------------------------------------------

func NewArrayList[E comparable]() *ArrayList[E] {
	return &ArrayList[E]{
		slice: []E{},
	}
}

func NewArrayListWithCap[E comparable](cap int) *ArrayList[E] {
	if cap < 0 {
		panic("invalid capacity")
	}
	return &ArrayList[E]{
		slice: make([]E, 0, cap),
	}
}

func NewArrayListWithEle[E comparable](elements ...E) *ArrayList[E] {
	list := NewArrayListWithCap[E](len(elements))
	for _, item := range elements {
		list.Add(item)
	}
	return list
}
