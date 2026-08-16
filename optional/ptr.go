package optional

// OfPtr returns an empty Optional for nil, otherwise an Optional containing
// a copy of the value pointed to by value.
func OfPtr[T any](value *T) Optional[T] {
	if value == nil {
		return Empty[T]()
	}
	return Of(*value)
}
