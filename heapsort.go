package heapsort

import . "sort"
import "fmt"
import "math"

// The heapsort package implements a simple heapsort sorting algorithm.
//
// HeapSort proceeds in two phases. First we build the binary tree with the invariant
// After that it's easy, we take the highest value located at the root, index 0
// put it at the end, and rebuild the tree with 1 less element
//
// The key data structure here, the heap, is really just a flattened complete binary tree.
// Complete means every level but the last is fully filled.
// Unlike many examples of Heapsort we use index 0.
// The heap is accessed as follows.
// Given an element e at location i:
// the left child node of e is at 2 * i + 1,
// the right child node of e is at 2 * i + 2,
// the parent node of e is at (i-1) / 2, which implies the root is on the "left", at index 0
// the invariant that must be satisifed is that s[parent(i)] >= s[i].

// lp prints a linear list of a heap of values
func Pl(d Interface, str string) {
    s := (d).(IntSlice)
    fmt.Printf("%s: ", str)
    for _, v := range s {
        fmt.Printf("%02d ", v)
    }
    fmt.Printf("\n");
    fmt.Printf("%s: ", str)
    for i := range s {
        fmt.Printf("%02d ", i)
    }
    fmt.Printf("\n")
}

func exp(x, y int) int {
  return int(math.Pow(float64(x), float64(y)))
}

// tp prints the tree of a heap of values
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

// traverse a (sub)tree downwards starting at r, but stopping when we get to the end, ci < ni.
// while traversing check the invariant and if not satisfied fix it by swapping the root and child
// this routine assumes that the nodes below it already satisfy the invariant
func h(d Interface, ni, ari int) {
    //s := (d).(IntSlice)
    //r := s[ari]

    ri := ari
    //si := ari
    //fmt.Printf("h(s, ni=%d, ri=%d), r=%d\n", ni, ri, r)
    for {
        ci := ri * 2 + 1
        if ci > ni {
            break
        }
        //fmt.Printf("ci=%d, ri=%d, ni=%d\n", ci, ri, ni)

        if ci < ni {
            //fmt.Printf("s[%d]=%d s[%d]=%d\n", ci, s[ci], ci+1, s[ci+1])
        }
        // follow the largest node
        if ci < ni && d.Less(ci, ci+1) { // s[ci] < s[ci+1]
            ci++
            //fmt.Printf("ci=%d\n", ci)
        }
        if ni < 1 || d.Less(ci, ri)  { // invariant holds // r >= s[ci]
            break
        }
        // invariant doesn't hold, copy child to root and descent to the next level of tree
        //fmt.Printf("descend: Swap(%d, %d) = (%d, %d)\n", ri, ci, s[ri], s[ci])
        d.Swap(ri, ci) // bubble up the child value, s[ri] has already been saved, and the new one is a copy //         s[ri] = s[ci] 
        //si = ci
        ri = ci
    }
    //fmt.Printf("exit: Swap(%d, %d) = (%d, %d)\n\n", ari, si, s[ari], s[si])
    //Pl(d, "b")
    //d.Swap(ari, si) // s[ri] = r
    //Pl(d, "a")
    //p(s)
}

// sort a slice using a heaport
func Heapsort(d Interface) {
    //s := (d).(IntSlice)
    //Pl(d, "s")
    //d.Swap(0, 1)
    //Pl(d, "e")
    //return
    n := d.Len()-1
    // build a heap with max element at index 0 by building a binary tree and repeatedly satisfying the invaraint from the bottom up
    for i := n/2; i >= 0; i-- {
        h(d, n, i)
        //Pl(d, "i")
/*
        if (i == 1) {
            return
        }
*/
    }
    //Pl(d, "m")
    //Pt(d)

    // repeatedly put the root (largest element) at the end of the slice and rebuild the heap
    for n > 0 {
        //fmt.Printf("Heapsort: Swap(%d, %d), Swap(%d, %d)\n", 0, n, s[0], s[n])
        d.Swap(0, n) // s[0], s[n] = s[n], s[0]
        n--
        h(d, n, 0)
        //Pl(d, "p")
        //Pt(d)
    }
}

