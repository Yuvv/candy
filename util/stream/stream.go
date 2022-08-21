package stream

import (
	joglang "github.com/yuvv/jog/lang"
	jogu "github.com/yuvv/jog/util"
	jogf "github.com/yuvv/jog/util/function"
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
	Supplier() jogf.Supplier[A]

	Accumulator() jogf.BiConsumer[A, T]

	Combiner() jogf.BinaryOperator[A]

	Finisher() jogf.Function[A, R]

	Characteristics() jogu.Set[Characteristics]
}

// BaseStream interface
type BaseStream[T any, S BaseStream[T, S]] interface {
	joglang.AutoCloseable

	Iterator()

	Spliterator() jogu.Spliterator[T]

	IsParallel() bool

	Sequential() S

	Parallel() S

	Unordered() S

	OnClose(runnable joglang.Runnable) S
}

// Stream interface
type Stream[T any] interface {
	BaseStream[T, Stream[T]]

	Filter(predicate jogf.Predicate[T]) Stream[T]

	Map[T, R](mapper jogf.Function[T, R]) Stream[R]

	FlatMap[T, R](mapper jogf.Function[T, Stream[R]]) Stream[R]

	Distinct() Stream[T]

	Sorted() Stream[T]
	SortedBy() Stream[T]

	Peek(consumer jogf.Consumer[T]) Stream[T]

	Skip(n int64) Stream[T]

	Limit(maxSize int64) Stream[T]

	ForEach(consumer jogf.Consumer[T])

	ForEachOrdered(consumer jogf.Consumer[T])

	ToArray() []T

	Reduce(accumulator jogf.BinaryOperator[T]) T

	ReduceWithIdentity(identity T, accumulator jogf.BinaryOperator[T]) T

	CollectBySupplier[T, R](supplier jogf.Supplier[R], consumer jogf.BiConsumer[R, T], biConsumer jogf.BiConsumer[R, R]) R

	Collect[T, A, R](collector Collector[T, A, R]) R

	Min(comparator jogu.Comparator[T]) jogu.Optional[T]

	Max(comparator jogu.Comparator[T]) jogu.Optional[T]

	Count() int64

	AnyMatch(predicate jogf.Predicate[T]) bool

	AllMatch(predicate jogf.Predicate[T]) bool

	NoneMatch(predicate jogf.Predicate[T]) bool

	FindFirst() jogu.Optional[T]

	FindAny() jogu.Optional[T]
}

func Of[T](values ...T) Stream[T] {
	// todo:
	return nil
}

func Empty[T]() Stream[T] {
	// todo:
	return nil
}

func Concat[T](a, b Stream[T]) Stream[T] {
	// todo:
	return nil
}
