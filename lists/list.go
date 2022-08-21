package lists

import "github.com/yuvv/candy/collections"

type List[T comparable] interface {
	collections.Collection[T]
}
