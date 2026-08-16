package collection

import (
	innerlist "container/list"
)

type _LinkedList[E any] struct {
	AbstractList[E]

	lst *innerlist.List
}

/// ------------------------------------------------------------------------------------

func NewLinkedList[E comparable]() *_LinkedList[E] {
	return &_LinkedList[E]{
		AbstractList: AbstractList[E]{
			equalMethod: func(e E, o any) bool {
				return e == o
			},
		},
		lst: innerlist.New(),
	}
}

func NewSpecLinkedList[E comparable](em func(e E, o any) bool) *_LinkedList[E] {
	return &_LinkedList[E]{
		AbstractList: AbstractList[E]{
			equalMethod: em,
		},
		lst: innerlist.New(),
	}
}
