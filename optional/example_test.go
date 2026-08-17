package optional_test

import (
	"fmt"

	"github.com/yuvv/candy/optional"
)

func Example() {
	name := optional.Of("candy")
	upper := optional.Map(name, func(value string) string {
		return "Hello, " + value
	})

	fmt.Println(upper.OrElse("Hello, world"))
	fmt.Println(optional.Empty[string]().OrElse("fallback"))
	// Output:
	// Hello, candy
	// fallback
}
