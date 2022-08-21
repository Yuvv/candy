package util

type Comparator[T any] interface {
	Compare(o1, o2 T) int

	Equals(obj any) bool
}
