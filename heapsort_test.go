package heapsort_test

import . "leb/heapsort"
//import "flag"
import "fmt"
import "math"
import "math/rand"
import "testing"


var r = rand.Float64

func lp(s []int, str string) {
    fmt.Printf("%s: ", str)
    for _, v := range s {
        fmt.Printf("%d ", v)
    }
    fmt.Printf("\n");
}

func exp(x, y int) int {
  return int(math.Pow(float64(x), float64(y)))
}

func p(s []int) {
    l := 0
    for i := range s {
        newlevel := exp(2, l) - 1
        if i == newlevel {
            l++
            fmt.Printf("\n%d: ", l)
        }
        fmt.Printf("%d ", s[i])
    }
    fmt.Printf("\n")
}

func rbetween(a int, b int) int {
        rf := r()
        diff := float64(b - a + 1)
        r2 := rf * diff
        r3 := r2 + float64(a)
        //fmt.Printf("rbetween: a=%d, b=%d, rf=%f, diff=%f, r2=%f, r3=%f\n", a, b, rf, diff, r2, r3)
        ret := int(r3)
        return ret
}

func verify(t *testing.T, s []int) {
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

func fill(s []int, a, b int) {
    for i := range s {
        s[i] = rbetween(a, b)
    }
}

func TestSimple(t *testing.T) {
    var s []int = []int{60, 94, 66, 44, 43, 68, 7, 16, 10, 30, 52, 81, 22, 38, 32}
    //var s []int = []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}

    //lp(s, "start")
    //p(s)
    Heapsort(s)
    //s[0] = 100
    //p(s)
    //lp(s, "finish")
    verify(t, s)
}

func BenchmarkBasic(b *testing.B) {
    b.StopTimer()
    s := make([]int, b.N)
    fill(s, 1, b.N*10)
    b.StartTimer()
    Heapsort(s)
    b.ReportAllocs()
}
