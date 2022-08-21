package lang

import (
	"github.com/yuvv/candy/util"
	"github.com/yuvv/candy/util/function"
)

type Iterable[T any] interface {
	Iterator() util.Iterator[T]

	ForEach(action function.Consumer[T])

	Spliterator() util.Spliterator[T]
}

//func (receiver Iterable[T]) ForEach(action function.Consumer[T]) {
//	if action == nil {
//		panic("action should not be nil")
//	}
//}
