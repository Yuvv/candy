package stream

import (
	"github.com/yuvv/candy/function"
	"github.com/yuvv/candy/iters"
	"github.com/yuvv/candy/lang"
	"github.com/yuvv/candy/optional"
	"github.com/yuvv/candy/sets"
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

	Characteristics() sets.Set[Characteristics]
}

// BaseStream interface
type BaseStream[T any, S BaseStream[T, S]] interface {
	lang.AutoCloseable

	Iterator()

	Spliterator() iters.Spliterator[T]

	IsParallel() bool

	Sequential() S

	Parallel() S

	Unordered() S

	OnClose(runnable lang.Runnable) S
}

// Stream interface
type Stream[T any] interface {
	BaseStream[T, Stream[T]]

	Filter(predicate function.Predicate[T]) Stream[T]

	// todo: @yuvv
	//Map[T, R any](mapper function.Function[T, R]) Stream[R]

	// todo: @yuvv
	//FlatMap[T, R any](mapper function.Function[T, Stream[R]]) Stream[R]

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

	// todo: @yuvv
	//CollectBySupplier[T, R any](supplier function.Supplier[R], consumer function.BiConsumer[R, T], biConsumer function.BiConsumer[R, R]) R

	// todo: @yuvv
	//Collect[T, A, R any](collector Collector[T, A, R]) R

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
	// todo:
	return nil
}

func Empty[T any]() Stream[T] {
	// todo:
	return nil
}

func Concat[T any](a, b Stream[T]) Stream[T] {
	// todo:
	return nil
}
