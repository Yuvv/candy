package arrays_test

import (
	"fmt"

	"github.com/yuvv/candy/arrays"
)

func Example() {
	even := arrays.Filter([]int{1, 2, 3, 4}, func(n int) bool {
		return n%2 == 0
	})
	doubled := arrays.Map(even, func(n int) int {
		return n * 2
	})

	fmt.Println(doubled)
	// Output: [4 8]
}
