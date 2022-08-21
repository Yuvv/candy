package collections

import (
	"github.com/yuvv/candy/lang"
	"github.com/yuvv/candy/util"
	"github.com/yuvv/candy/util/function"
)

type Collection[E any] interface {
	lang.Iterable[E]

	Size() int

	IsEmpty() bool

	Contains(o any) bool

	ToArray() []E

	Add(e E) bool

	Remove(o any) bool

	ContainsAll(collection Collection[E]) bool

	AddAll(collection Collection[E]) bool

	RemoveAll(collection Collection[E]) bool

	RemoveIf(predicate function.Predicate[E]) bool

	RetainAll(collection Collection[E]) bool

	Clear()

	//Stream() stream.Stream[E]
	//
	//ParallelStream() stream.Stream[E]
}

// AbstractCollection is the abstraction of Collection
type AbstractCollection[E any] struct {
}

func (a *AbstractCollection[E]) Iterator() util.Iterator[E] {
	//TODO implement me
	panic("implement me")
}

func (a *AbstractCollection[E]) ForEach(action function.Consumer[E]) {
	it := a.Iterator()
	for it.HasNext() {
		action(it.Next())
	}
}

func (a *AbstractCollection[E]) Spliterator() util.Spliterator[E] {
	//TODO implement me
	panic("implement me")
}

func (a *AbstractCollection[E]) Size() int {
	//TODO implement me
	panic("implement me")
}

func (a *AbstractCollection[E]) IsEmpty() bool {
	return a.Size() == 0
}

func (a *AbstractCollection[E]) Contains(o any) bool {
	//TODO implement me
	panic("implement me")
}

func (a *AbstractCollection[E]) ToArray() []E {
	res := make([]E, 0, a.Size())
	it := a.Iterator()
	for it.HasNext() {
		res = append(res, it.Next())
	}
	return res
}

func (a *AbstractCollection[E]) Add(e E) bool {
	//TODO implement me
	panic("implement me")
}

func (a *AbstractCollection[E]) Remove(o any) bool {
	//TODO implement me
	panic("implement me")
}

func (a *AbstractCollection[E]) ContainsAll(collection Collection[E]) bool {
	//TODO implement me
	panic("implement me")
}

func (a *AbstractCollection[E]) AddAll(collection Collection[E]) bool {
	//TODO implement me
	panic("implement me")
}

func (a *AbstractCollection[E]) RemoveAll(collection Collection[E]) bool {
	//TODO implement me
	panic("implement me")
}

func (a *AbstractCollection[E]) RemoveIf(predicate function.Predicate[E]) bool {
	//TODO implement me
	panic("implement me")
}

func (a *AbstractCollection[E]) RetainAll(collection Collection[E]) bool {
	//TODO implement me
	panic("implement me")
}

func (a *AbstractCollection[E]) Clear() {
	//TODO implement me
	panic("implement me")
}
