package zsort

// Partition rearranges the elements of x so that all elements for which pred
// returns true are placed before all elements for which pred returns false.
// There is no guarantee about the order of the elements within each partition.
// Partition returns the index of the first element in the second partition.
//
// Note that the predicate receives the index of the element in the slice, not
// the element itself.
func Partition[S ~[]E, E any](x S, pred func(int) bool) int {
	split := 0
	front := 0
	back := len(x) - 1

Outer:
	for ; front < back; front++ {
		if !pred(front) {
			for ; back > front; back-- {
				if pred(back) {
					x[front], x[back] = x[back], x[front]
					break
				}
			}
			if back <= front {
				break Outer
			}
		}
		split++
	}

	return split
}
