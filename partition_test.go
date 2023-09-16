package zsort_test

import (
	"slices"
	"testing"

	"github.com/arl/zsort"
)

func TestPartition(t *testing.T) {
	s := []int{7, 1, 1, 7, 1, 1, 7}
	split := zsort.Partition(s, func(i int) bool { return s[i] >= 3 })
	if split != 3 {
		t.Errorf("Partition() = %v, want %v", split, 3)
	}

	want := []int{7, 7, 7, 1, 1, 1, 1}
	if !slices.Equal(s, want) {
		t.Errorf("Partition: slice = %+v, want %+v", s, want)
	}
}

func TestPartitionInts(t *testing.T) {
	tests := []struct {
		data  []int
		pred  func([]int, int) bool // pred receives the slice since we can't refer to 'data' from 'pred'
		split int
		want  []int
	}{
		{
			data: []int{7, 1, 1, 9, 1, 1, 3}, pred: func(s []int, i int) bool { return s[i] >= 3 },
			split: 3, want: []int{7, 3, 9, 1, 1, 1, 1},
		},
		{
			data: []int{7, 3, 9, 1, 1, 1, 1}, pred: func(s []int, i int) bool { return s[i] == 1 },
			split: 4, want: []int{1, 1, 1, 1, 9, 3, 7},
		},
		{
			data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, pred: func(s []int, i int) bool { return s[i]%3 == 0 },
			split: 3, want: []int{9, 6, 3, 4, 5, 2, 7, 8, 1},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			split := zsort.Partition(tt.data, func(i int) bool { return tt.pred(tt.data, i) })
			if split != tt.split {
				t.Errorf("Partition() = %v, want %v", split, 1)
			}

			if !slices.Equal(tt.data, tt.want) {
				t.Errorf("Partition: slice = %+v, want %+v", tt.data, tt.want)
			}
		})
	}
}
