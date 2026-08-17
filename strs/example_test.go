package strs_test

import (
	"fmt"

	"github.com/yuvv/candy/strs"
)

func ExampleNewIntJoiner() {
	joined := strs.NewIntJoiner(", ").
		WithPrefix("[").
		WithSuffix("]").
		Join([]int{1, 2, 3})

	fmt.Println(joined)
	// Output: [1, 2, 3]
}
