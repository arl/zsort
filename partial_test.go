package zsort_test

import (
	"fmt"
	"math/rand"
	"reflect"
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

func TestPartialNilSlice(t *testing.T) {
	var s []int
	zsort.Partial(s, 10)
}

func TestPartialPrecise(t *testing.T) {
	// Precise in the sense we verify the whole slice is as expected after
	// applying Partial to it (as opposed to TestPartial where we only test that
	// the first 10% are sorted).
	org := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	for n := 0; n < 12; n++ {
		s := make([]int, len(org))
		copy(s, org)
		zsort.Partial(s, n)

		want := make([]int, len(org))
		for i := range want {
			if i < n {
				want[i] = i
			} else {
				want[i] = 10 - i + n
			}
		}

		if !reflect.DeepEqual(s, want) {
			t.Errorf("for n = %d, got = %v, want = %v", n, s, want)
		}
	}
}

func BenchmarkPartialInt1K(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0x2cc
		}
		b.StartTimer()
		zsort.Partial(data, 10)
		b.StopTimer()
	}
}

func BenchmarkPartialInt1K_Sorted(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = i
		}
		b.StartTimer()
		zsort.Partial(data, 10)
		b.StopTimer()
	}
}

func BenchmarkPartialInt1K_Reversed(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = len(data) - i
		}
		b.StartTimer()
		zsort.Partial(data, 10)
		b.StopTimer()
	}
}

func BenchmarkPartialInt1K_Mod8(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = i % 8
		}
		b.StartTimer()
		zsort.Partial(data, 10)
		b.StopTimer()
	}
}
