package util

import (
	"github.com/yuvv/candy/lang"
	"github.com/yuvv/candy/util/function"
)

type Optional[T any] struct {
	// if non-null, the value;
	// if null, indicates no value is present.
	value T
}

func (receiver *Optional[T]) Get() T {
	if receiver.value == nil {
		panic("No value present")
	}
	return receiver.value
}

func (receiver *Optional[T]) IsPresent() bool {
	return receiver.value != nil
}

func (receiver *Optional[T]) IsEmpty() bool {
	return receiver.value == nil
}

func (receiver *Optional[T]) IfPresent(consumer function.Consumer[T]) {
	if receiver.value != nil {
		consumer(receiver.value)
	}
}

func (receiver *Optional[T]) IfPresentOrElse(consumer function.Consumer[T], emptyAction lang.Runnable) {
	if receiver.value != nil {
		consumer(receiver.value)
	} else {
		emptyAction.Run()
	}
}

func Empty[T any]() *Optional[T] {
	return &Optional[T]{}
}

func Of[T any](t T) *Optional[T] {
	return &Optional[T]{value: t}
}
