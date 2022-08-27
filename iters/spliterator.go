package iters

import (
	"github.com/yuvv/candy/function"
	"github.com/yuvv/candy/lang"
)

type Spliterator[T any] interface {
	TryAdvance(comparator lang.Comparator[T]) bool

	ForEachRemaining(consumer function.Consumer[T])

	TrySplit() Spliterator[T]

	EstimateSize() int64

	GetExactSizeIfKnown() int64

	Characteristics() int8

	HasCharacteristics(characteristics int8) bool

	GetComparator() lang.Comparator[T]
}
