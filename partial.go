package zsort

import (
	"cmp"
)

// void
// __partial_sort(_RandomAccessIterator __first, __middle, __last,
// 	_Compare __comp)
// __comp should return true if the first element is less than the second

func Partial[S ~[]E, E cmp.Ordered](x S, n int) {
	if n == 0 || len(x) == 0 {
		return
	}
	// __make_heap<_Compare>(__first, __middle, __comp);
	makeHeap(x[:n])
	// typename iterator_traits<_RandomAccessIterator>::difference_type __len = __middle - __first;
	// for (_RandomAccessIterator __i = __middle; __i != __last; ++__i)
	// {
	for i := n; i < len(x); i++ {
		//     if (__comp(*__i, *__first))
		//     {
		if x[i] < x[0] {
			//         swap(*__i, *__first);
			x[i], x[0] = x[0], x[i]
			//         __sift_down<_Compare>(__first, __middle, __comp, __len, __first);
			siftDown(x[:n], n, 0)
		}
		//     }
		// }
	}

	// Measure:
	//  either
	// slices.Sort(x[:n])
	//  or:
	// sortHeap(x[:n])

	// __sort_heap<_Compare>(__first, __middle, __comp);
	// TODO
	// fmt.Println(x[:n])

	// sortHeap(x[:n])

	// This is the inlined version of sortHeap(x[:n])
	for i := n; i > 1; n, i = n-1, i-1 {
		if i > 1 {
			x[0], x[n-1] = x[n-1], x[0]
			siftDown(x[:n-1], i-1, 0)
		}
	}
}

// void
// __make_heap(_RandomAccessIterator __first, __last, _Compare __comp)
func makeHeap[S ~[]E, E cmp.Ordered](x S) {
	// typedef typename iterator_traits<_RandomAccessIterator>::difference_type difference_type;
	// difference_type __n = __last - __first;
	n := len(x)
	if n > 1 {
		// start from the first parent, there is no need to consider children
		// for (difference_type __start = (__n - 2) / 2; __start >= 0; --__start)
		// {
		//     __sift_down<_Compare>(__first, __last, __comp, __n, __first + __start);
		// }
		for start := (n - 2) / 2; start >= 0; start-- {
			siftDown(x, n, start)
		}
	}
}

// void
// __sort_heap(_RandomAccessIterator __first, _RandomAccessIterator __last, _Compare __comp)
func sortHeap[S ~[]E, E cmp.Ordered](x S) {
	// typedef typename iterator_traits<_RandomAccessIterator>::difference_type difference_type;
	// for (difference_type __n = __last - __first; __n > 1; --__last, --__n)
	//     __pop_heap<_Compare>(__first, __last, __comp, __n);
	last := len(x)
	for n := last; n > 1; last, n = last-1, n-1 {
		popHeap(x[:last], n)
	}
}

// __pop_heap(_RandomAccessIterator __first, _RandomAccessIterator __last, _Compare __comp,
//
//	typename iterator_traits<_RandomAccessIterator>::difference_type __len)
func popHeap[S ~[]E, E cmp.Ordered](x S, n int) {
	// if (__len > 1)
	// {
	//     swap(*__first, *--__last);
	//     __sift_down<_Compare>(__first, __last, __comp, __len - 1, __first);
	// }
	if n > 1 {
		x[0], x[len(x)-1] = x[len(x)-1], x[0]
		siftDown(x[:len(x)-1], n-1, 0)
	}
}

// void
// __sift_down(_RandomAccessIterator __first, _RandomAccessIterator /*__last*/,
//
//	_Compare __comp,
//	typename iterator_traits<_RandomAccessIterator>::difference_type __len,
//	_RandomAccessIterator __start)
func siftDown[S ~[]E, E cmp.Ordered](x S, n, start int) {
	// left-child of __start is at 2 * __start + 1
	// right-child of __start is at 2 * __start + 2
	//    difference_type __child = __start - __first;
	child := start

	//    if (__len < 2 || (__len - 2) / 2 < __child)
	//        return;
	if n < 2 || (n-2)/2 < child {
		return
	}

	//    __child = 2 * __child + 1;
	child = 2*child + 1
	//    _RandomAccessIterator __child_i = __first + __child;
	ichild := child

	// if ((__child + 1) < __len && __comp(*__child_i, *(__child_i + 1))) {
	//     // right-child exists and is greater than left-child
	//     ++__child_i;
	//     ++__child;
	// }
	if (child+1) < n && x[ichild] < x[ichild+1] {
		// Right-child exists and is greater than left-child.
		ichild++
		child++
	}

	// // check if we are in heap-order
	// if (__comp(*__child_i, *__start))
	// // we are, __start is larger than it's largest child
	// return;

	// Check if we are in heap-order
	if x[ichild] < x[start] {
		// We are, x[0] is larger than its largest child.
		return
	}

	// value_type __top(_VSTD::move(*__start));
	top := x[start]

	// do
	// {
	for {
		// // we are not in heap-order, swap the parent with it's largest child
		// *__start = _VSTD::move(*__child_i);
		// __start = __child_i;

		// We are not in heap-order, swap the parent with its largest child.
		x[start] = x[ichild]
		start = ichild

		// if ((__len - 2) / 2 < __child)
		//     break;
		if (n-2)/2 < child {
			break
		}

		// // recompute the child based off of the updated parent
		// __child = 2 * __child + 1;
		// __child_i = __first + __child;

		// Recompute the child based off of the updated parent.
		child = 2*child + 1
		ichild = child

		// if ((__child + 1) < __len && __comp(*__child_i, *(__child_i + 1))) {
		//     // right-child exists and is greater than left-child
		//     ++__child_i;
		//     ++__child;
		// }

		if (child+1) < n && x[ichild] < x[ichild+1] {
			// Right-child exists and is greater than left-child.
			ichild++
			child++
		}

		// // check if we are in heap-order
		// } while (!__comp(*__child_i, __top));

		// TODO(arl) improve stop condition readability

		// check if we are in heap-order
		if x[ichild] >= top {
			continue
		}
		break
	}

	// *__start = _VSTD::move(__top);
	x[start] = top
}
