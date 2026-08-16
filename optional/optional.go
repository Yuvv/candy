package optional

// Optional is a container object which may or may not contain a value.
type Optional[T any] struct {
	value   T
	present bool
}

// Empty returns an empty Optional.
func Empty[T any]() Optional[T] {
	return Optional[T]{}
}

// Of returns an Optional describing value.
func Of[T any](value T) Optional[T] {
	return Optional[T]{
		value:   value,
		present: true,
	}
}

// IsPresent returns true if there is a value present.
func (receiver Optional[T]) IsPresent() bool {
	return receiver.present
}

// IsEmpty returns true if there is no value present.
func (receiver Optional[T]) IsEmpty() bool {
	return !receiver.present
}

// Get returns the present value or panics if no value is present.
func (receiver Optional[T]) Get() T {
	if receiver.IsEmpty() {
		panic("optional: no value present")
	}
	return receiver.value
}

// OrElse returns the present value, otherwise defaultValue.
func (receiver Optional[T]) OrElse(defaultValue T) T {
	if receiver.IsPresent() {
		return receiver.value
	}
	return defaultValue
}

// OrElseGet returns the present value, otherwise the value supplied by supplier.
func (receiver Optional[T]) OrElseGet(supplier func() T) T {
	if receiver.IsPresent() {
		return receiver.value
	}
	return supplier()
}

// IfPresent calls consumer with the value if there is one present.
func (receiver Optional[T]) IfPresent(consumer func(T)) {
	if receiver.IsPresent() {
		consumer(receiver.value)
	}
}

// Map returns an Optional containing the mapped value if optional is present.
func Map[T any, R any](optional Optional[T], mapper func(T) R) Optional[R] {
	if optional.IsEmpty() {
		return Empty[R]()
	}
	return Of(mapper(optional.value))
}

// FlatMap returns the Optional produced by mapper if optional is present.
func FlatMap[T any, R any](optional Optional[T], mapper func(T) Optional[R]) Optional[R] {
	if optional.IsEmpty() {
		return Empty[R]()
	}
	return mapper(optional.value)
}
