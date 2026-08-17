package stream

import (
	"github.com/yuvv/candy/function"
	"github.com/yuvv/candy/lang"
	"github.com/yuvv/candy/optional"
)

type sequentialStream[T any] struct {
	values []T
}

func newSequentialStream[T any](values []T) Stream[T] {
	copied := make([]T, len(values))
	copy(copied, values)
	return sequentialStream[T]{values: copied}
}

func (stream sequentialStream[T]) Filter(predicate function.Predicate[T]) Stream[T] {
	panic("stream: Filter is not implemented")
}

func (stream sequentialStream[T]) Distinct() Stream[T] {
	panic("stream: Distinct is not implemented")
}

func (stream sequentialStream[T]) Sorted() Stream[T] {
	panic("stream: Sorted is not implemented")
}

func (stream sequentialStream[T]) SortedBy() Stream[T] {
	panic("stream: SortedBy is not implemented")
}

func (stream sequentialStream[T]) Peek(consumer function.Consumer[T]) Stream[T] {
	panic("stream: Peek is not implemented")
}

func (stream sequentialStream[T]) Skip(n int64) Stream[T] {
	panic("stream: Skip is not implemented")
}

func (stream sequentialStream[T]) Limit(maxSize int64) Stream[T] {
	panic("stream: Limit is not implemented")
}

func (stream sequentialStream[T]) ForEach(consumer function.Consumer[T]) {
	for _, value := range stream.values {
		consumer(value)
	}
}

func (stream sequentialStream[T]) ForEachOrdered(consumer function.Consumer[T]) {
	stream.ForEach(consumer)
}

func (stream sequentialStream[T]) ToArray() []T {
	values := make([]T, len(stream.values))
	copy(values, stream.values)
	return values
}

func (stream sequentialStream[T]) Reduce(accumulator function.BinaryOperator[T]) T {
	if len(stream.values) == 0 {
		var zero T
		return zero
	}

	result := stream.values[0]
	for _, value := range stream.values[1:] {
		result = accumulator(result, value)
	}
	return result
}

func (stream sequentialStream[T]) ReduceWithIdentity(identity T, accumulator function.BinaryOperator[T]) T {
	result := identity
	for _, value := range stream.values {
		result = accumulator(result, value)
	}
	return result
}

func (stream sequentialStream[T]) Min(comparator lang.Comparator[T]) optional.Optional[T] {
	if len(stream.values) == 0 {
		return optional.Empty[T]()
	}

	minimum := stream.values[0]
	for _, value := range stream.values[1:] {
		if comparator.Compare(value, minimum) < 0 {
			minimum = value
		}
	}
	return optional.Of(minimum)
}

func (stream sequentialStream[T]) Max(comparator lang.Comparator[T]) optional.Optional[T] {
	if len(stream.values) == 0 {
		return optional.Empty[T]()
	}

	maximum := stream.values[0]
	for _, value := range stream.values[1:] {
		if comparator.Compare(value, maximum) > 0 {
			maximum = value
		}
	}
	return optional.Of(maximum)
}

func (stream sequentialStream[T]) Count() int64 {
	return int64(len(stream.values))
}

func (stream sequentialStream[T]) AnyMatch(predicate function.Predicate[T]) bool {
	panic("stream: AnyMatch is not implemented")
}

func (stream sequentialStream[T]) AllMatch(predicate function.Predicate[T]) bool {
	panic("stream: AllMatch is not implemented")
}

func (stream sequentialStream[T]) NoneMatch(predicate function.Predicate[T]) bool {
	panic("stream: NoneMatch is not implemented")
}

func (stream sequentialStream[T]) FindFirst() optional.Optional[T] {
	if len(stream.values) == 0 {
		return optional.Empty[T]()
	}
	return optional.Of(stream.values[0])
}

func (stream sequentialStream[T]) FindAny() optional.Optional[T] {
	return stream.FindFirst()
}
