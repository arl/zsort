package zsort

import "cmp"

// Partial sorts the first n elements of x, such that they are the smallest n
// elements in ascending order, using the type's natural ordering. The remaining
// elements in the slice are not sorted but are guaranteed to be greater than or
// equal to the elements in the first n positions.
//
// x is sorted in-place. The sort is not guaranteed to be stable: equal elements
// may be reversed from their original order
func Partial[S ~[]E, E cmp.Ordered](x S, n int) {
	if n == 0 || len(x) == 0 {
		return
	}

	makeHeap(x[:n])

	for i := n; i < len(x); i++ {
		if x[i] < x[0] {
			x[i], x[0] = x[0], x[i]
			siftDown(x[:n], n, 0)
		}
	}

	for i := n; i > 1; n, i = n-1, i-1 {
		if i > 1 {
			x[0], x[n-1] = x[n-1], x[0]
			siftDown(x[:n-1], i-1, 0)
		}
	}
}

func makeHeap[S ~[]E, E cmp.Ordered](x S) {
	n := len(x)
	if n > 1 {
		// Start from the first parent, there is no need to consider children.
		for start := (n - 2) / 2; start >= 0; start-- {
			siftDown(x, n, start)
		}
	}
}

func sortHeap[S ~[]E, E cmp.Ordered](x S) {
	last := len(x)
	for n := last; n > 1; last, n = last-1, n-1 {
		popHeap(x[:last], n)
	}
}

func popHeap[S ~[]E, E cmp.Ordered](x S, n int) {
	if n > 1 {
		x[0], x[len(x)-1] = x[len(x)-1], x[0]
		siftDown(x[:len(x)-1], n-1, 0)
	}
}

func siftDown[S ~[]E, E cmp.Ordered](x S, n, start int) {
	// left-child of __start is at 2 * __start + 1
	// right-child of __start is at 2 * __start + 2
	child := start

	if n < 2 || (n-2)/2 < child {
		return
	}

	child = 2*child + 1
	ichild := child

	if (child+1) < n && x[ichild] < x[ichild+1] {
		// Right-child exists and is greater than left-child.
		ichild++
		child++
	}

	// Check if we are in heap-order.
	if x[ichild] < x[start] {
		// We are, x[0] is larger than its largest child.
		return
	}

	top := x[start]

	for {
		// We are not in heap-order, swap the parent with its largest child.
		x[start] = x[ichild]
		start = ichild

		if (n-2)/2 < child {
			break
		}

		// Recompute the child based off of the updated parent.
		child = 2*child + 1
		ichild = child

		if (child+1) < n && x[ichild] < x[ichild+1] {
			// Right-child exists and is greater than left-child.
			ichild++
			child++
		}

		// Check if we are in heap-order
		if x[ichild] < top {
			break
		}
	}

	x[start] = top
}
