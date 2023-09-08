package zsort_test

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"

	"github.com/arl/zsort"
)

func TestPartial(t *testing.T) {
	sizes := []int{0, 10, 100, 1000, 10000}
	for _, size := range sizes {
		t.Run(fmt.Sprintf("size=%d", size), func(t *testing.T) {
			n := size / 10
			s := make([]int, 0, size)
			for i := 0; i < size; i++ {
				s = append(s, rand.Intn(10000))
			}

			zsort.Partial(s, n)
			if !sort.IntsAreSorted(s[:n]) {
				t.Errorf("s[:%d] should be sorted", n)
			}

			if n > 0 {
				if sort.IntsAreSorted(s[n:]) {
					t.Errorf("s[:%d] should not be sorted", n)
				}
			}
		})
	}
}
