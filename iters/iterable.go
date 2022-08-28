package iters

import (
	"github.com/yuvv/candy/function"
)

type Iterable[T comparable] interface {
	Iterator() Iterator[T]

	ForEach(action function.Consumer[T])

	Spliterator() Spliterator[T]
}

//func (receiver Iterable[T]) ForEach(action function.Consumer[T]) {
//	if action == nil {
//		panic("action should not be nil")
//	}
//}
