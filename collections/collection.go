package collections

import (
	"github.com/yuvv/candy/function"
	"github.com/yuvv/candy/iters"
)

type Collection[E comparable] interface {
	iters.Iterable[E]

	Size() int

	IsEmpty() bool

	Contains(o E) bool

	ToArray() []E

	Add(e E) bool

	Remove(o E) bool

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
type AbstractCollection[E comparable] struct {
}

// Iterator is abstract method
func (a *AbstractCollection[E]) Iterator() iters.Iterator[E] {
	panic("implement me")
}

func (a *AbstractCollection[E]) ForEach(action function.Consumer[E]) {
	it := a.Iterator()
	for it.HasNext() {
		action(it.Next())
	}
}

func (a *AbstractCollection[E]) Spliterator() iters.Spliterator[E] {
	//TODO implement me
	panic("implement me")
}

// Size is abstract method
func (a *AbstractCollection[E]) Size() int {
	panic("implement me")
}

func (a *AbstractCollection[E]) IsEmpty() bool {
	return a.Size() == 0
}

func (a *AbstractCollection[E]) Contains(o E) bool {
	it := a.Iterator()
	for it.HasNext() {
		if it.Next() == o {
			return true
		}
	}
	return false
}

func (a *AbstractCollection[E]) ToArray() []E {
	res := make([]E, 0, a.Size())
	it := a.Iterator()
	for it.HasNext() {
		res = append(res, it.Next())
	}
	return res
}

// Add is abstract method
func (a *AbstractCollection[E]) Add(e E) bool {
	panic("implement me")
}

// Remove will remove the first element equals o
func (a *AbstractCollection[E]) Remove(o E) bool {
	it := a.Iterator()
	for it.HasNext() {
		if it.Next() == o {
			it.Remove()
			return true
		}
	}
	return true
}

func (a *AbstractCollection[E]) ContainsAll(collection Collection[E]) bool {
	it := collection.Iterator()
	for it.HasNext() {
		if !a.Contains(it.Next()) {
			return false
		}
	}
	return true
}

func (a *AbstractCollection[E]) AddAll(collection Collection[E]) bool {
	modified := false
	it := collection.Iterator()
	for it.HasNext() {
		modified = a.Add(it.Next())
	}
	return modified
}

func (a *AbstractCollection[E]) RemoveAll(collection Collection[E]) bool {
	modified := false
	it := a.Iterator()
	for it.HasNext() {
		if collection.Contains(it.Next()) {
			it.Remove()
			modified = true
		}
	}
	return modified
}

func (a *AbstractCollection[E]) RemoveIf(predicate function.Predicate[E]) bool {
	//TODO implement me
	panic("implement me")
}

func (a *AbstractCollection[E]) RetainAll(collection Collection[E]) bool {
	modified := false
	it := a.Iterator()
	for it.HasNext() {
		if !collection.Contains(it.Next()) {
			it.Remove()
			modified = true
		}
	}
	return modified
}

func (a *AbstractCollection[E]) Clear() {
	it := a.Iterator()
	for it.HasNext() {
		_ = it.Next()
		it.Remove()
	}
}
