package zsort_test

import (
	"fmt"
	"testing"

	"github.com/arl/zsort"
)

// func TestPartial(t *testing.T) {
// 	s := []int{5, 7, 4, 2, 8, 6, 1, 9, 0, 3}
// 	fmt.Println(s)

// 	zsort.PartialFunc(s, 3, cmp.Less)
// 	fmt.Println(s)
// 	zsort.PartialFunc(s, 4, cmp.Less)
// 	fmt.Println(s)
// 	zsort.PartialFunc(s, len(s), cmp.Less)
// 	fmt.Println(s)
// }

func TestPartial(t *testing.T) {
	s := []int{5, 7, 4, 2, 8, 6, 1, 9, 0, 3}
	fmt.Println(s)

	zsort.Partial(s, 3)
	fmt.Println("zsort.Partial(s, 3)")
	fmt.Println(s)

	zsort.Partial(s, 4)
	fmt.Println("zsort.Partial(s, 4)")
	fmt.Println(s)

	zsort.Partial(s, 8)
	fmt.Println("zsort.Partial(s, 8)")
	fmt.Println(s)

	zsort.Partial(s, len(s))
	fmt.Println("zsort.Partial(s, len(s))")
	fmt.Println(s)
}

// std::array<int, 10> s {5, 7, 4, 2, 8, 6, 1, 9, 0, 3};
// print(s, 0);
// std::partial_sort(s.begin(), s.begin() + 3, s.end());
// print(s, 3);
// std::partial_sort(s.rbegin(), s.rbegin() + 4, s.rend());
// print(s, -4);
// std::partial_sort(s.rbegin(), s.rbegin() + 5, s.rend(), std::greater{});
// print(s, -5);
