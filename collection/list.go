package collection

import (
	"github.com/yuvv/candy/function"
	"github.com/yuvv/candy/iter"
	"github.com/yuvv/candy/lang"
)

type List[E any] interface {
	Collection[E]

	AddAllAt(idx int, collection Collection[E])

	ReplaceAll(operator function.UnaryOperator[E])

	Sort(comparator lang.Comparator[E])

	Get(idx int) E

	Set(idx int, ele E) E

	AddAt(idx int, ele E)

	RemoveAt(idx int) E

	IndexOf(o E) int

	LastIndexOf(o E) int

	ListIterator() iter.ListIterator[E]

	ListIteratorFrom(idx int) iter.ListIterator[E]

	SubList(fromIdx, toIdx int) List[E]

	getModCount() int
}

type AbstractList[E any] struct {
	AbstractCollection[E]

	modCount int

	equalMethod func(x E, other any) bool
}

func (a *AbstractList[E]) Iterator() iter.Iterator[E] {
	panic("collection: AbstractList.Iterator must be implemented by concrete list")
}

func (a *AbstractList[E]) ForEach(action function.Consumer[E]) {
	panic("collection: AbstractList.ForEach must be implemented by concrete list")
}

func (a *AbstractList[E]) Spliterator() iter.Spliterator[E] {
	panic("collection: AbstractList.Spliterator must be implemented by concrete list")
}

func (a *AbstractList[E]) Contains(o E) bool {
	panic("collection: AbstractList.Contains must be implemented by concrete list")
}

func (a *AbstractList[E]) ToArray() []E {
	panic("collection: AbstractList.ToArray must be implemented by concrete list")
}

func (a *AbstractList[E]) Add(e E) bool {
	a.AddAt(a.Size(), e)
	return true
}

func (a *AbstractList[E]) Remove(o E) bool {
	panic("collection: AbstractList.Remove must be implemented by concrete list")
}

func (a *AbstractList[E]) ContainsAll(collection Collection[E]) bool {
	panic("collection: AbstractList.ContainsAll must be implemented by concrete list")
}

func (a *AbstractList[E]) AddAll(collection Collection[E]) bool {
	panic("collection: AbstractList.AddAll must be implemented by concrete list")
}

func (a *AbstractList[E]) RemoveAll(collection Collection[E]) bool {
	panic("collection: AbstractList.RemoveAll must be implemented by concrete list")
}

func (a *AbstractList[E]) RemoveIf(predicate function.Predicate[E]) bool {
	panic("collection: AbstractList.RemoveIf must be implemented by concrete list")
}

func (a *AbstractList[E]) RetainAll(collection Collection[E]) bool {
	panic("collection: AbstractList.RetainAll must be implemented by concrete list")
}

func (a *AbstractList[E]) AddAllAt(idx int, collection Collection[E]) {
	panic("collection: AbstractList.AddAllAt must be implemented by concrete list")
}

func (a *AbstractList[E]) ReplaceAll(operator function.UnaryOperator[E]) {
	panic("collection: AbstractList.ReplaceAll must be implemented by concrete list")
}

func (a *AbstractList[E]) Sort(comparator lang.Comparator[E]) {
	panic("collection: AbstractList.Sort must be implemented by concrete list")
}

func (a *AbstractList[E]) Get(idx int) E {
	panic("collection: AbstractList.Get must be implemented by concrete list")
}

func (a *AbstractList[E]) Set(idx int, ele E) {
	panic("collection: AbstractList.Set must be implemented by concrete list")
}

func (a *AbstractList[E]) AddAt(idx int, ele E) {
	panic("collection: AbstractList.AddAt must be implemented by concrete list")
}

func (a *AbstractList[E]) RemoveAt(idx int) E {
	panic("collection: AbstractList.RemoveAt must be implemented by concrete list")
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

func (a *AbstractList[E]) ListIterator() iter.ListIterator[E] {
	panic("collection: AbstractList.ListIterator must be implemented by concrete list")
}

func (a *AbstractList[E]) ListIteratorFrom(idx int) iter.ListIterator[E] {
	panic("collection: AbstractList.ListIteratorFrom must be implemented by concrete list")
}

func (a *AbstractList[E]) SubList(fromIdx, toIdx int) List[E] {
	panic("collection: AbstractList.SubList must be implemented by concrete list")
}

func (a *AbstractList[E]) GetEleEqualMethod() func(x E, other any) bool {
	return a.equalMethod
}
