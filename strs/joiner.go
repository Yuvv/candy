package strs

import "strconv"

func JoinSlice[T any](items []T, adapt func(T) string, prefix, delimiter, suffix string) string {
	b := NewBuilder().AppendString(prefix)
	for i, item := range items {
		if i > 0 {
			b.AppendString(delimiter)
		}
		b.AppendString(adapt(item))
	}
	b.AppendString(suffix)
	return b.String()
}

type Joiner[T any] struct {
	prefix    string
	delimiter string
	suffix    string
	adapt     func(T) string
}

func NewJoiner[T any](delimiter string, adapt func(T) string) *Joiner[T] {
	return &Joiner[T]{delimiter: delimiter, adapt: adapt}
}

func (j *Joiner[T]) WithPrefix(prefix string) *Joiner[T] {
	j.prefix = prefix
	return j
}

func (j *Joiner[T]) WithSuffix(suffix string) *Joiner[T] {
	j.suffix = suffix
	return j
}

func (j *Joiner[T]) Join(items []T) string {
	return JoinSlice(items, j.adapt, j.prefix, j.delimiter, j.suffix)
}

func NewStringJoiner(delimiter string) *Joiner[string] {
	return NewJoiner(delimiter, func(value string) string { return value })
}

func NewIntJoiner(delimiter string) *Joiner[int] {
	return NewJoiner(delimiter, strconv.Itoa)
}

func NewInt32Joiner(delimiter string) *Joiner[int32] {
	return NewJoiner(delimiter, func(value int32) string { return strconv.FormatInt(int64(value), 10) })
}

func NewInt64Joiner(delimiter string) *Joiner[int64] {
	return NewJoiner(delimiter, func(value int64) string { return strconv.FormatInt(value, 10) })
}

func NewUintJoiner(delimiter string) *Joiner[uint] {
	return NewJoiner(delimiter, func(value uint) string { return strconv.FormatUint(uint64(value), 10) })
}

func NewUint64Joiner(delimiter string) *Joiner[uint64] {
	return NewJoiner(delimiter, func(value uint64) string { return strconv.FormatUint(value, 10) })
}
