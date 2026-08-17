package stream

import "sort"

// Map applies mapper to every element in source order.
func Map[T any, R any](s Stream[T], mapper func(T) R) Stream[R] {
	if s == nil {
		return Empty[R]()
	}

	source := s.ToArray()
	values := make([]R, 0, len(source))
	for _, value := range source {
		values = append(values, mapper(value))
	}
	return newSequentialStream(values)
}

// FlatMap applies mapper to every element in source order and concatenates the resulting streams.
func FlatMap[T any, R any](s Stream[T], mapper func(T) Stream[R]) Stream[R] {
	if s == nil {
		return Empty[R]()
	}

	var values []R
	for _, value := range s.ToArray() {
		mapped := mapper(value)
		if mapped != nil {
			values = append(values, mapped.ToArray()...)
		}
	}
	return newSequentialStream(values)
}

// SortedBy returns a sorted copy of s using less.
func SortedBy[T any](s Stream[T], less func(a, b T) bool) Stream[T] {
	if less == nil {
		panic("stream: SortedBy requires a less function")
	}
	if s == nil {
		return Empty[T]()
	}

	source := s.ToArray()
	values := make([]T, len(source))
	copy(values, source)
	sort.SliceStable(values, func(i, j int) bool {
		return less(values[i], values[j])
	})
	return newSequentialStream(values)
}
