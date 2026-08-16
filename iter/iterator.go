package iter

import (
	"github.com/yuvv/candy/function"
)

type Iterator[E any] interface {
	HasNext() bool

	Next() E

	Remove()

	ForEachRemaining(action function.Consumer[E])
}

type AbstractIterator[E any] struct {
}

func (a *AbstractIterator[E]) HasNext() bool {
	panic("iter: AbstractIterator.HasNext must be implemented by concrete iterator")
}

func (a *AbstractIterator[E]) Next() E {
	panic("iter: AbstractIterator.Next must be implemented by concrete iterator")
}

func (a *AbstractIterator[E]) Remove() {
	panic("iter: AbstractIterator.Remove must be implemented by concrete iterator")
}

func (a *AbstractIterator[E]) ForEachRemaining(action function.Consumer[E]) {
	for a.HasNext() {
		action(a.Next())
	}
}

type ListIterator[E any] interface {
	Iterator[E]

	HasPrevious() bool

	Previous() E

	NextIndex() int

	PreviousIndex() int

	Set(e E)

	Add(e E)
}
