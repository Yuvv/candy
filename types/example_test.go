package types_test

import (
	"fmt"

	"github.com/yuvv/candy/types"
)

func ExampleBitFlag() {
	const (
		read uint8 = 1 << iota
		write
		execute
	)

	var permissions types.BitFlag[uint8]
	permissions.Set(read | write)
	permissions.Clear(write)

	fmt.Println(permissions.Has(read), permissions.Has(write), permissions.Has(execute))
	// Output: true false false
}

func ExampleBitmap() {
	var bitmap types.Bitmap
	bitmap.Set(0)
	bitmap.Set(3)

	fmt.Println(bitmap.Get(0), bitmap.Get(1), bitmap.Get(3))
	fmt.Println(bitmap.String())
	// Output:
	// true false true
	// b:1001
}
