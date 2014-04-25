package heapsort
// The heapsort package implements a simple heapsort sorting algorithm.

// The key data structure here, the heap, is really just a flattened complete binary tree.
// Complete means every level but the last is fully filled.
// The heap is accessed as follows.
// Given an element e at location i:
// the left child node of e is at 2 * i,
// the right child node of e is at 2 * i + 1,
// the parent node of e is at i / 2, which implies the root is on the "left", ie at index 0
// the invariant that must be satisifed is that s[parent(i)] >= s[i].

// HeapSort proceeds in two phases. First we build the binary tree with the invariant
// After that it's easy, we take the highest value located at the root, index 0
// put it at the end, and rebuild the tree with 1 less element



// traverse a (sub)tree downwards starting at r, but stopping when we get to the end, ci < ni.
// while traversing check the invariant and if not satisfied fix it by swapping the root and child
// this routine assumes that the nodes below it already satisfy the invariant
func h(s []int, ni, ri int) {
    r := s[ri]
    //fmt.Printf("h(s, ni=%d, ri=%d), r=%d\n", ni, ri, r)
    for ri <= ni/2 {
        ci := ri * 2
        //fmt.Printf("b=%d, r=%d, n=%d\n", ci, ri, ni)

        if ci < ni {
            //fmt.Printf("s[%d]=%d s[%d]=%d\n", ci, s[ci], ci+1, s[ci+1])
        }
        // follow the largest node
        if ci < ni && s[ci] < s[ci+1] {
            ci++
            //fmt.Printf("ci=%d\n", ci)
        }
        if r >= s[ci] { // invariant holds
            break
        }
        // invariant doesn't hold, copy child to root and descent to the next level of tree
        //fmt.Printf("descend: s[%d]=%d = s[%d]=%d\n", ri, s[ri], ci, s[ci])
        s[ri] = s[ci] // bubble up the child value, s[ri] has already been saved, and the new one is a copy
        ri = ci
    }
    //fmt.Printf("exit: s[%d]=%d\n", ri, r)
    s[ri] = r
    //p(s)
}


// sort a slice using a heaport
func Heapsort(s []int) {
    n := len(s)-1
    // build a heap with max element at index 0 by building a binary tree and repeatedly satisfying the invaraint from the bottom up
    for i := n/2; i >= 0; i-- {
        h(s, n, i)
    }
    //p(s)

    // repeatedly put the root (largest element) at the end of the slice and rebuild the heap
    for n > 0 {
        //fmt.Printf("Swap(%d, %d), Swap(%d, %d)\n", 0, n, s[0], s[n])
        s[0], s[n] = s[n], s[0]
        n--
        h(s, n, 0)
    }
}

