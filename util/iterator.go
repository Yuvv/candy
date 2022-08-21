package util

import "github.com/yuvv/candy/util/function"

type Iterator[E any] interface {
	HasNext() bool

	Next() E

	Remove()

	ForEachRemaining(action function.Consumer[E])
}

//func (receiver Iterator[E]) ForEachRemaining(action function.Consumer[E]) {
//	if action == nil {
//		panic("Consumer should not be nil")
//	}
//	for receiver.HasNext() {
//		action(receiver.Next())
//	}
//}
