// Copyright © 2014 Lawrence E. Bakst. All rights reserved.
package heapsort_test

import . "leb/heapsort"
//import "flag"
import "fmt"
import "time"
import "math/rand"
import "testing"
import . "sort"


var r = rand.Float64

func rbetween(a int, b int) int {
        rf := r()
        diff := float64(b - a + 1)
        r2 := rf * diff
        r3 := r2 + float64(a)
        //fmt.Printf("rbetween: a=%d, b=%d, rf=%f, diff=%f, r2=%f, r3=%f\n", a, b, rf, diff, r2, r3)
        ret := int(r3)
        return ret
}

// verify that a slice is in order
func verify(t *testing.T, s IntSlice) {
    pv := -1
    for k, v := range s {
        if k == 0 {
            pv = v
            continue
        }
        if pv > v {
            t.Error("Sort out of order")
            return
        }
    }
}

// fill a slice with randome numbers
func fill(s IntSlice, a, b int) {
    for i := range s {
        s[i] = rbetween(a, b)
    }
}

func TestTrivial(t *testing.T) {
    var s IntSlice = []int{60, 94, 66, 44, 43, 68, 7, 16, 10, 30, 52, 81, 22, 38, 32}
    //var s []int = []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}

    //pl(s, "start")
    //pt(s)
    Heapsort(s)
    //s[0] = 100
    //pl(s)
    Pl(s, "e")
    Pt(s)
    verify(t, s)
}


func TestBasic(t *testing.T) {
    s := make(IntSlice, 15)
    rand.Seed(1)
    fill(s, 1, 99)
    Heapsort(s)
    verify(t, s)
}


func TestAdvanced(t *testing.T) {
    n := 100
    rand.Seed(2)
    s := make(IntSlice, n)
    fill(s, 1, n*3)
    Heapsort(s)
    verify(t, s)
}

func TestExtended(t *testing.T) {
    for i := 1; i <= 100; i++ { // start with 1 to get the seed we're used to
        //t.Logf("i=%d\n", i)
        seed := int64(0)
        // fixed pattern or different values each time
        if false {
            seed = int64(i)
        } else {
            seed = time.Now().UTC().UnixNano()
        }
        rand.Seed(seed)
        n := rbetween(10, 100000)
        fmt.Printf("%d: %d ", i, n)
        s := make(IntSlice, n)
        fill(s, 1, n*3)
        Heapsort(s)
        verify(t, s)
    }
    fmt.Printf("\n")
}

/*
func TestSpecific(t *testing.T) {
        i := int64(11)
        t.Logf("i=%d\n", i)
        rand.Seed(i)
        n := rbetween(10, 100)
        t.Logf("n=%d\n", n)
        s := make(IntSlice, n)
        fill(s, 1, n*3)
        Heapsort(s)
        verify(t, s)
}
*/

func BenchmarkBasic(b *testing.B) {
    b.StopTimer()
    s := make(IntSlice, b.N)
    fill(s, 1, b.N*10)
    b.StartTimer()
    Heapsort(s)
    b.ReportAllocs()
}
