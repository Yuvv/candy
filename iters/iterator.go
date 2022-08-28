package iters

import (
	"github.com/yuvv/candy/function"
)

type Iterator[E comparable] interface {
	HasNext() bool

	Next() E

	Remove()

	ForEachRemaining(action function.Consumer[E])
}

type AbstractIterator[E comparable] struct {
}

func (a *AbstractIterator[E]) HasNext() bool {
	panic("implement me")
}

func (a *AbstractIterator[E]) Next() E {
	panic("implement me")
}

func (a *AbstractIterator[E]) Remove() {
	panic("implement me")
}

func (a *AbstractIterator[E]) ForEachRemaining(action function.Consumer[E]) {
	for a.HasNext() {
		action(a.Next())
	}
}

type ListIterator[E comparable] interface {
	Iterator[E]

	HasPrevious() bool

	Previous() E

	NextIndex() int

	PreviousIndex() int

	Set(e E)

	Add(e E)
}
