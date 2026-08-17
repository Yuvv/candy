package stream

import (
	"github.com/yuvv/candy/collection"
	"github.com/yuvv/candy/function"
	"github.com/yuvv/candy/lang"
	"github.com/yuvv/candy/optional"
)

type Characteristics int8

const (
	C_CONCURRENT Characteristics = iota
	C_UNORDERED
	C_IDENTITY_FINISH
)

type StreamShape int8

const (
	SS_REFERENCE StreamShape = iota
	SS_INT_VALUE
	SS_LONG_VALUE
	SS_DOUBLE_VALUE
)

type Collector[T, A, R any] interface {
	Supplier() function.Supplier[A]

	Accumulator() function.BiConsumer[A, T]

	Combiner() function.BinaryOperator[A]

	Finisher() function.Function[A, R]

	Characteristics() collection.Set[Characteristics]
}

// BaseStream interface
//type BaseStream[T any, S BaseStream[T, S]] interface {
//	lang.AutoCloseable
//
//	Iterator()
//
//	Spliterator() iter.Spliterator[T]
//
//	IsParallel() bool
//
//	Sequential() S
//
//	Parallel() S
//
//	Unordered() S
//
//	OnClose(runnable lang.Runnable) S
//}

// Stream interface
type Stream[T any] interface {
	// Go 1.18 does not support generic methods, so map and flat-map style
	// operations are not represented here.

	Filter(predicate function.Predicate[T]) Stream[T]

	Distinct() Stream[T]

	Sorted() Stream[T]
	SortedBy() Stream[T]

	Peek(consumer function.Consumer[T]) Stream[T]

	Skip(n int64) Stream[T]

	Limit(maxSize int64) Stream[T]

	ForEach(consumer function.Consumer[T])

	ForEachOrdered(consumer function.Consumer[T])

	ToArray() []T

	Reduce(accumulator function.BinaryOperator[T]) T

	ReduceWithIdentity(identity T, accumulator function.BinaryOperator[T]) T

	Min(comparator lang.Comparator[T]) optional.Optional[T]

	Max(comparator lang.Comparator[T]) optional.Optional[T]

	Count() int64

	AnyMatch(predicate function.Predicate[T]) bool

	AllMatch(predicate function.Predicate[T]) bool

	NoneMatch(predicate function.Predicate[T]) bool

	FindFirst() optional.Optional[T]

	FindAny() optional.Optional[T]
}

// -------------------------------------------------------------------------------------------------------------------

// Of func
func Of[T any](values ...T) Stream[T] {
	return newSequentialStream(values)
}

func Empty[T any]() Stream[T] {
	return newSequentialStream([]T(nil))
}

func Concat[T any](a, b Stream[T]) Stream[T] {
	var values []T
	if a != nil {
		values = append(values, a.ToArray()...)
	}
	if b != nil {
		values = append(values, b.ToArray()...)
	}
	return newSequentialStream(values)
}
