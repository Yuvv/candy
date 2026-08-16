package collection

import (
	"github.com/yuvv/candy/function"
	"github.com/yuvv/candy/iter"
)

type Collection[E any] interface {
	iter.Iterable[E]

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
}

// AbstractCollection is the abstraction of Collection
type AbstractCollection[E any] struct {

	// GetEleEqualMethod return the func to check elements equals
	GetEleEqualMethod func() func(x E, other any) bool

	iteratorMethod func() iter.Iterator[E]
	sizeMethod     func() int
	addMethod      func(E) bool
}

// Iterator is abstract method
func (a *AbstractCollection[E]) Iterator() iter.Iterator[E] {
	if a.iteratorMethod == nil {
		panic("collection: AbstractCollection.Iterator must be implemented by concrete collection")
	}
	return a.iteratorMethod()
}

func (a *AbstractCollection[E]) ForEach(action function.Consumer[E]) {
	it := a.Iterator()
	for it.HasNext() {
		action(it.Next())
	}
}

func (a *AbstractCollection[E]) Spliterator() iter.Spliterator[E] {
	panic("collection: AbstractCollection.Spliterator is unsupported")
}

// Size is abstract method
func (a *AbstractCollection[E]) Size() int {
	if a.sizeMethod == nil {
		panic("collection: AbstractCollection.Size must be implemented by concrete collection")
	}
	return a.sizeMethod()
}

func (a *AbstractCollection[E]) IsEmpty() bool {
	return a.Size() == 0
}

func (a *AbstractCollection[E]) Contains(o E) bool {
	it := a.Iterator()
	for it.HasNext() {
		if a.GetEleEqualMethod()(it.Next(), o) {
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
	if a.addMethod == nil {
		panic("collection: AbstractCollection.Add must be implemented by concrete collection")
	}
	return a.addMethod(e)
}

// Remove will remove the first element equals o
func (a *AbstractCollection[E]) Remove(o E) bool {
	it := a.Iterator()
	for it.HasNext() {
		if a.GetEleEqualMethod()(it.Next(), o) {
			it.Remove()
			return true
		}
	}
	return false
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
		modified = a.Add(it.Next()) || modified
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
	modified := false
	it := a.Iterator()
	for it.HasNext() {
		if predicate(it.Next()) {
			it.Remove()
			modified = true
		}
	}
	return modified
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
