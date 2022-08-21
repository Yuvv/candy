package sets

import "github.com/yuvv/candy/collections"

type Set[T comparable] interface {
	collections.Collection[T]
}
