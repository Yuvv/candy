package lists

import (
	"github.com/yuvv/candy/collections"
	"github.com/yuvv/candy/util"
	"github.com/yuvv/candy/util/function"
)

type ArrayList[T any] struct {
	slice []T
}

func (lst *ArrayList[T]) Iterator() util.Iterator[T] {
	//TODO implement me
	panic("implement me")
}

func (lst *ArrayList[T]) ForEach(action function.Consumer[T]) {
	//TODO implement me
	panic("implement me")
}

func (lst *ArrayList[T]) Spliterator() util.Spliterator[T] {
	//TODO implement me
	panic("implement me")
}

func (lst *ArrayList[T]) Size() int {
	//TODO implement me
	panic("implement me")
}

func (lst *ArrayList[T]) IsEmpty() bool {
	//TODO implement me
	panic("implement me")
}

func (lst *ArrayList[T]) Contains(o any) bool {
	//TODO implement me
	panic("implement me")
}

func (lst *ArrayList[T]) ToArray() []T {
	//TODO implement me
	panic("implement me")
}

func (lst *ArrayList[T]) Add(e T) bool {
	//TODO implement me
	panic("implement me")
}

func (lst *ArrayList[T]) Remove(o any) bool {
	//TODO implement me
	panic("implement me")
}

func (lst *ArrayList[T]) ContainsAll(collection collections.Collection[T]) bool {
	//TODO implement me
	panic("implement me")
}

func (lst *ArrayList[T]) AddAll(collection collections.Collection[T]) bool {
	//TODO implement me
	panic("implement me")
}

func (lst *ArrayList[T]) RemoveAll(collection collections.Collection[T]) bool {
	//TODO implement me
	panic("implement me")
}

func (lst *ArrayList[T]) RemoveIf(predicate function.Predicate[T]) bool {
	//TODO implement me
	panic("implement me")
}

func (lst *ArrayList[T]) RetainAll(collection collections.Collection[T]) bool {
	//TODO implement me
	panic("implement me")
}

func (lst *ArrayList[T]) Clear() {
	//TODO implement me
	panic("implement me")
}

/// ---------------------------------------------------------------------------

func NewArrayList[T any]() *ArrayList[T] {
	return &ArrayList[T]{
		slice: []T{},
	}
}

func NewArrayListWithCap[T any](cap int) *ArrayList[T] {
	if cap < 0 {
		panic("invalid capacity")
	}
	return &ArrayList[T]{
		slice: make([]T, 0, cap),
	}
}
