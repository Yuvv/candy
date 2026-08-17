package concurrent_test

import (
	"context"
	"fmt"

	"github.com/yuvv/candy/concurrent"
)

func ExampleAwaitTask() {
	results, err := concurrent.AwaitTask(
		context.Background(),
		func(_ context.Context, n int) (int, error) {
			return n * n, nil
		},
		[]int{1, 2, 3},
	)

	fmt.Println(results, err == nil)
	// Output: [1 4 9] true
}
