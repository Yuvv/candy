package stream

import (
	jogu "github.com/yuvv/jog/util"
	"github.com/yuvv/jog/util/function"
)

func SupportBySpliterator[T](spliterator jogu.Spliterator[T], parallel bool) Stream[T] {
	// todo:
	return nil
}

func SupportBySupplier[T](supplier function.Supplier[jogu.Spliterator[T]], characteristics Characteristics, parallel bool) Stream[T] {
	// todo:
	return nil
}
