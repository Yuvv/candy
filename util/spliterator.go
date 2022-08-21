package util

import "github.com/yuvv/candy/util/function"

type Spliterator[T any] interface {
	TryAdvance(comparator Comparator[T]) bool

	ForEachRemaining(consumer function.Consumer[T])

	TrySplit() Spliterator[T]

	EstimateSize() int64

	GetExactSizeIfKnown() int64

	Characteristics() int8

	HasCharacteristics(characteristics int8) bool

	GetComparator() Comparator[T]
}
