package lists

import (
	"github.com/yuvv/candy/collections"
	"github.com/yuvv/candy/function"
	"github.com/yuvv/candy/iters"
	"github.com/yuvv/candy/lang"
)

type List[E any] interface {
	collections.Collection[E]

	AddAllAt(idx int, collection collections.Collection[E])

	ReplaceAll(operator function.UnaryOperator[E])

	Sort(comparator lang.Comparator[E])

	Get(idx int) E

	Set(idx int, ele E) E

	AddAt(idx int, ele E)

	RemoveAt(idx int) E

	IndexOf(o E) int

	LastIndexOf(o E) int

	ListIterator() iters.ListIterator[E]

	ListIteratorFrom(idx int) iters.ListIterator[E]

	SubList(fromIdx, toIdx int) List[E]

	getModCount() int
}

type AbstractList[E any] struct {
	collections.AbstractCollection[E]

	modCount int

	equalMethod func(x E, other any) bool
}

func (a *AbstractList[E]) Iterator() iters.Iterator[E] {
	//TODO implement me
	panic("implement me")
}

func (a *AbstractList[E]) ForEach(action function.Consumer[E]) {
	//TODO implement me
	panic("implement me")
}

func (a *AbstractList[E]) Spliterator() iters.Spliterator[E] {
	//TODO implement me
	panic("implement me")
}

func (a *AbstractList[E]) Contains(o E) bool {
	//TODO implement me
	panic("implement me")
}

func (a *AbstractList[E]) ToArray() []E {
	//TODO implement me
	panic("implement me")
}

func (a *AbstractList[E]) Add(e E) bool {
	a.AddAt(a.Size(), e)
	return true
}

func (a *AbstractList[E]) Remove(o E) bool {
	//TODO implement me
	panic("implement me")
}

func (a *AbstractList[E]) ContainsAll(collection collections.Collection[E]) bool {
	//TODO implement me
	panic("implement me")
}

func (a *AbstractList[E]) AddAll(collection collections.Collection[E]) bool {
	//TODO implement me
	panic("implement me")
}

func (a *AbstractList[E]) RemoveAll(collection collections.Collection[E]) bool {
	//TODO implement me
	panic("implement me")
}

func (a *AbstractList[E]) RemoveIf(predicate function.Predicate[E]) bool {
	//TODO implement me
	panic("implement me")
}

func (a *AbstractList[E]) RetainAll(collection collections.Collection[E]) bool {
	//TODO implement me
	panic("implement me")
}

func (a *AbstractList[E]) AddAllAt(idx int, collection collections.Collection[E]) {
	//TODO implement me
	panic("implement me")
}

func (a *AbstractList[E]) ReplaceAll(operator function.UnaryOperator[E]) {
	//TODO implement me
	panic("implement me")
}

func (a *AbstractList[E]) Sort(comparator lang.Comparator[E]) {
	//TODO implement me
	panic("implement me")
}

func (a *AbstractList[E]) Get(idx int) E {
	//TODO implement me
	panic("implement me")
}

func (a *AbstractList[E]) Set(idx int, ele E) {
	//TODO implement me
	panic("implement me")
}

func (a *AbstractList[E]) AddAt(idx int, ele E) {
	//TODO implement me
	panic("implement me")
}

func (a *AbstractList[E]) RemoveAt(idx int) E {
	//TODO implement me
	panic("implement me")
}

func (a *AbstractList[E]) IndexOf(o E) int {
	it := a.ListIterator()
	for it.HasNext() {
		if a.equalMethod(it.Next(), o) {
			return it.PreviousIndex()
		}
	}
	return -1
}

func (a *AbstractList[E]) LastIndexOf(o E) int {
	it := a.ListIteratorFrom(a.Size())
	for it.HasNext() {
		if a.equalMethod(it.Next(), o) {
			return it.NextIndex()
		}
	}
	return -1
}

func (a *AbstractList[E]) ListIterator() iters.ListIterator[E] {
	//TODO implement me
	panic("implement me")
}

func (a *AbstractList[E]) ListIteratorFrom(idx int) iters.ListIterator[E] {
	//TODO implement me
	panic("implement me")
}

func (a *AbstractList[E]) SubList(fromIdx, toIdx int) List[E] {
	//TODO implement me
	panic("implement me")
}

func (a *AbstractList[E]) GetEleEqualMethod() func(x E, other any) bool {
	return a.equalMethod
}
