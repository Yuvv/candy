package stream

import (
	"github.com/yuvv/candy/function"
	jogu "github.com/yuvv/candy/iters"
)

func SupportBySpliterator[T any](spliterator jogu.Spliterator[T], parallel bool) Stream[T] {
	// todo:
	return nil
}

func SupportBySupplier[T any](supplier function.Supplier[jogu.Spliterator[T]], characteristics Characteristics, parallel bool) Stream[T] {
	// todo:
	return nil
}
