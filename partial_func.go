package zsort

// PartialFunc sorts the first n elements of x, such that they are the smallest
// n elements, given the provided less function. The remaining elements in the
// slice are not sorted but are guaranteed to be greater than or equal to the
// elements in the first n positions.
//
// x is sorted in-place. The sort is not guaranteed to be stable: equal elements
// may be reversed from their original order
func PartialFunc[S ~[]E, E any](x S, n int, less func(E, E) bool) {
	if n == 0 || len(x) == 0 {
		return
	}

	makeHeap_func(x[:n], less)

	for i := n; i < len(x); i++ {
		if less(x[i], x[0]) {
			x[i], x[0] = x[0], x[i]
			siftDown_func(x[:n], n, 0, less)
		}
	}

	for i := n; i > 1; n, i = n-1, i-1 {
		if i > 1 {
			x[0], x[n-1] = x[n-1], x[0]
			siftDown_func(x[:n-1], i-1, 0, less)
		}
	}
}

func makeHeap_func[S ~[]E, E any](x S, less func(E, E) bool) {
	n := len(x)
	if n > 1 {
		// Start from the first parent, there is no need to consider children.
		for start := (n - 2) / 2; start >= 0; start-- {
			siftDown_func(x, n, start, less)
		}
	}
}

func sortHeap_func[S ~[]E, E any](x S, less func(E, E) bool) {
	last := len(x)
	for n := last; n > 1; last, n = last-1, n-1 {
		popHeap_func(x[:last], n, less)
	}
}

func popHeap_func[S ~[]E, E any](x S, n int, less func(E, E) bool) {
	if n > 1 {
		x[0], x[len(x)-1] = x[len(x)-1], x[0]
		siftDown_func(x[:len(x)-1], n-1, 0, less)
	}
}

func siftDown_func[S ~[]E, E any](x S, n, start int, less func(E, E) bool) {
	// left-child of __start is at 2 * __start + 1
	// right-child of __start is at 2 * __start + 2
	child := start

	if n < 2 || (n-2)/2 < child {
		return
	}

	child = 2*child + 1
	ichild := child

	if (child+1) < n && less(x[ichild], x[ichild+1]) {
		// Right-child exists and is greater than left-child.
		ichild++
		child++
	}

	// Check if we are in heap-order
	if less(x[ichild], x[start]) {
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

		if (child+1) < n && less(x[ichild], x[ichild+1]) {
			// Right-child exists and is greater than left-child.
			ichild++
			child++
		}

		// Check if we are in heap-order
		if less(x[ichild], top) {
			break
		}
	}

	x[start] = top
}
