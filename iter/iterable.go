package iter

import (
	"github.com/yuvv/candy/function"
)

type Iterable[T any] interface {
	Iterator() Iterator[T]

	ForEach(action function.Consumer[T])

	Spliterator() Spliterator[T]
}
