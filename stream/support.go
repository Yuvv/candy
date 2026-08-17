package stream

import (
	"github.com/yuvv/candy/function"
	jogu "github.com/yuvv/candy/iter"
)

func SupportBySpliterator[T any](spliterator jogu.Spliterator[T], parallel bool) Stream[T] {
	panic("stream: SupportBySpliterator is not implemented")
}

func SupportBySupplier[T any](supplier function.Supplier[jogu.Spliterator[T]], characteristics Characteristics, parallel bool) Stream[T] {
	panic("stream: SupportBySupplier is not implemented")
}
