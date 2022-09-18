package sets

import "github.com/yuvv/candy/collections"

type Set[T any] interface {
	collections.Collection[T]
}
