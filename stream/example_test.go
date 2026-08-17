package stream_test

import (
	"fmt"
	"strconv"

	"github.com/yuvv/candy/stream"
)

func ExampleOf() {
	values := stream.Of(1, 2, 3).
		Filter(func(value int) bool { return value%2 == 1 }).
		ToArray()

	fmt.Println(values)
	// Output: [1 3]
}

func ExampleMap() {
	values := stream.Map(stream.Of(1, 2, 3), strconv.Itoa).ToArray()

	fmt.Println(values)
	// Output: [1 2 3]
}

func ExampleSortedBy() {
	type person struct {
		name string
		age  int
	}

	people := stream.Of(
		person{name: "Grace", age: 37},
		person{name: "Ada", age: 31},
		person{name: "Linus", age: 34},
	)
	ordered := stream.SortedBy(people, func(a, b person) bool {
		return a.age < b.age
	}).ToArray()

	for _, person := range ordered {
		fmt.Println(person.name, person.age)
	}
	// Output:
	// Ada 31
	// Linus 34
	// Grace 37
}
