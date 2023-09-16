package zsort_test

import (
	"fmt"

	"github.com/arl/zsort"
)

func ExamplePartial() {
	s := []int{5, 2, 6, 3, 1, 4, 0, 9, 8, 7}

	zsort.Partial(s, 3)

	fmt.Println(s[:3])
	// Output: [0 1 2]
}

func ExamplePartialFunc() {
	type person struct {
		name string
		age  int
	}

	persons := []person{
		{"Alice", 19},
		{"Bob", 17},
		{"Jane", 36},
		{"Henri", 24},
		{"Sandra", 22},
	}

	zsort.PartialFunc(persons, 3, func(p1, p2 person) bool { return p1.age > p2.age })

	fmt.Println(persons[:3])
	// Output: [{Jane 36} {Henri 24} {Sandra 22}]
}

func ExamplePartition() {
	s := []int{7, 1, 1, 7, 1, 1, 7}

	split := zsort.Partition(s, func(i int) bool { return s[i] >= 3 })

	fmt.Println("split =", split)
	fmt.Println(s)

	// Output:
	// split = 3
	// [7 7 7 1 1 1 1]
}
