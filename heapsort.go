// Copyright © 2014 Lawrence E. Bakst. All rights reserved.
// The heapsort package implements a simple heapsort algorithm.
//
// HeapSort proceeds in two phases. First we build a complete binary tree called the heap, while maintaining an invariant.
// Second, we take the highest value, now located at the root, index 0, and swap it with the last element.
// After the swap rebuild the heap with 1 less element, as the largest elements are repeatedly swapped to the end.
//
// The key data structure here, the heap, is really just a flattened complete binary tree.
// Complete means every level but the last is fully filled.
// Unlike many examples of Heapsort we use index 0, most start at index 1 
// The heap is accessed as follows.
// Given an element e at location i:
// the left child node of e is at 2 * i + 1,
// the right child node of e is at 2 * i + 2,
// the parent node of e is at (i-1) / 2, which implies the root is on the "left", at index 0
// the invariant that must be satisifed is that s[parent(i)] >= s[i].
package heapsort

import . "sort"
import "fmt"
import "math"

// Pl prints a linear list of a heap of values and indicies underneath them.
func Pl(d Interface, str string) {
	s := (d).(IntSlice)
	fmt.Printf("%s: ", str)
	for _, v := range s {
		fmt.Printf("%02d ", v)
	}
	fmt.Printf("\n")
	fmt.Printf("%s: ", str)
	for i := range s {
		fmt.Printf("%02d ", i)
	}
	fmt.Printf("\n")
}

func exp(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

// Pt prints the heap as a formatted binary tree.
func Pt(d Interface) {
	s := (d).(IntSlice)
	l := 0
	for i := range s {
		newlevel := exp(2, l) - 1
		if i == newlevel {
			l++
			fmt.Printf("\n%d: ", l)
		}
		fmt.Printf("%d ", s[i])
	}
	fmt.Printf("\n\n")
}

// Knuth calles this function "siftup" and others call this "heapify". I hate both, but Knuth wins.
// Traverse a (sub)tree downwards starting at r, but stopping when we get to the end, ci < ni.
// While traversing check the invariant and if not satisfied fix it by swapping the root and child.
// NB: This function assumes that the nodes at the levels below the root already satisfy the invariant.
// EG: In order to build a healp from scratch all nodes at level l-1 must satisfy the invariant before level l.
func siftup(d Interface, ni, ri int) {
	//s := (d).(IntSlice)
	//r := s[ari]

	//    ri := ari
	//fmt.Printf("h(s, ni=%d, ri=%d), r=%d\n", ni, ri, r)
	for ri < (ni+1)/2 {
		ci := ri*2 + 1 // calculate left child node
		//fmt.Printf("ci=%d, ri=%d, ni=%d\n", ci, ri, ni)

		if ci < ni && d.Less(ci, ci+1) { // follow the largest node left or right
			ci++
			//fmt.Printf("ci=%d\n", ci)
		}
		if d.Less(ci, ri) { // invariant holds // ni < 1 || 
			break
		}
		// invariant doesn't hold, swap child and root and descent to the next level of tree
		//fmt.Printf("descend: Swap(%d, %d) = (%d, %d)\n", ri, ci, s[ri], s[ci])
		d.Swap(ri, ci)
		ri = ci
	}
}

// Heapsort sorts data using a heaport algorithm. It implements the standard Go sort.Interface interface.
func Heapsort(d Interface) {
	n := d.Len() - 1
	// build a heap with max element at index 0 by building a binary tree and repeatedly satisfying the invaraint from the bottom up
	for i := n / 2; i >= 0; i-- {
		siftup(d, n, i)
		//Pl(d, "i")
	}
	//Pl(d, "m")
	//Pt(d)

	// repeatedly put the root (largest element) at the end of the slice and rebuild the heap with one less element
	for n > 0 {
		//fmt.Printf("Heapsort: Swap(%d, %d), Swap(%d, %d)\n", 0, n, s[0], s[n])
		d.Swap(0, n)
		n--
		siftup(d, n, 0)
		//Pl(d, "p")
		//Pt(d)
	}
}

