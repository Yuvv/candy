package stream

import (
	"math"
	"reflect"
	"sort"

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
	values := make([]T, 0, len(stream.values))
	for _, value := range stream.values {
		if predicate(value) {
			values = append(values, value)
		}
	}
	return newSequentialStream(values)
}

func (stream sequentialStream[T]) Distinct() Stream[T] {
	values := make([]T, 0, len(stream.values))
	for _, value := range stream.values {
		distinct := true
		for _, existing := range values {
			if reflect.DeepEqual(value, existing) {
				distinct = false
				break
			}
		}
		if distinct {
			values = append(values, value)
		}
	}
	return newSequentialStream(values)
}

func (stream sequentialStream[T]) Sorted() Stream[T] {
	kind := reflect.TypeOf((*T)(nil)).Elem().Kind()
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64, reflect.String:
	default:
		panic("stream: Sorted requires an ordered element type")
	}

	values := stream.ToArray()
	sort.SliceStable(values, func(i, j int) bool {
		left := reflect.ValueOf(values[i])
		right := reflect.ValueOf(values[j])
		switch kind {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return left.Int() < right.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return left.Uint() < right.Uint()
		case reflect.Float32, reflect.Float64:
			leftFloat := left.Float()
			rightFloat := right.Float()
			leftNaN := math.IsNaN(leftFloat)
			rightNaN := math.IsNaN(rightFloat)
			if leftNaN {
				return false
			}
			if rightNaN {
				return true
			}
			return leftFloat < rightFloat
		case reflect.String:
			return left.String() < right.String()
		}
		return false
	})
	return newSequentialStream(values)
}

func (stream sequentialStream[T]) SortedBy() Stream[T] {
	panic("stream: SortedBy requires a comparator; use stream.SortedBy")
}

func (stream sequentialStream[T]) Peek(consumer function.Consumer[T]) Stream[T] {
	for _, value := range stream.values {
		consumer(value)
	}
	return newSequentialStream(stream.values)
}

func (stream sequentialStream[T]) Skip(n int64) Stream[T] {
	if n <= 0 {
		return newSequentialStream(stream.values)
	}
	if n >= int64(len(stream.values)) {
		return Empty[T]()
	}
	return newSequentialStream(stream.values[int(n):])
}

func (stream sequentialStream[T]) Limit(maxSize int64) Stream[T] {
	if maxSize < 0 {
		panic("stream: limit size must be non-negative")
	}
	if maxSize == 0 {
		return Empty[T]()
	}
	if maxSize >= int64(len(stream.values)) {
		return newSequentialStream(stream.values)
	}
	return newSequentialStream(stream.values[:int(maxSize)])
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
	for _, value := range stream.values {
		if predicate(value) {
			return true
		}
	}
	return false
}

func (stream sequentialStream[T]) AllMatch(predicate function.Predicate[T]) bool {
	for _, value := range stream.values {
		if !predicate(value) {
			return false
		}
	}
	return true
}

func (stream sequentialStream[T]) NoneMatch(predicate function.Predicate[T]) bool {
	for _, value := range stream.values {
		if predicate(value) {
			return false
		}
	}
	return true
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
