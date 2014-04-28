// Copyright © 2014 Lawrence E. Bakst. All rights reserved.
package main
import "leb/heapsort"
import "flag"
import "time"
import "fmt"
import "os"
import "io"
import "bufio"
import "runtime"
import "errors"
import . "sort"

// simple porgram to read a file, sort it, write it

var mf = flag.Bool("m", false, "dump memory stats")
var tf = flag.Bool("t", false, "show timing")
var vf = flag.Bool("v", false, "verbose - show read and write cnts")
var memstats *runtime.MemStats

func dump_mstats(m *runtime.MemStats, cstats bool, gc bool) {
    fmt.Printf("Alloc=%d, TotalAlloc=%d, Sys=%d, Lookups=%d, Mallocs=%d, Frees=%d\n", m.Alloc, m.TotalAlloc, m.Sys, m.Lookups, m.Mallocs, m.Frees)
    fmt.Printf("HeapAlloc=%d, HeapSys=%d, HeapIdle=%d, HeapInuse=%d, HeapReleased=%d, HeapObjects=%d\n", m.HeapAlloc, m.HeapSys, m.HeapIdle, m.HeapInuse, m.HeapReleased, m.HeapObjects)
    fmt.Printf("StackInuse=%d, StackSys=%d, MSpanInuse=%d, MSpanSys=%d, MCacheSys=%d, BuckHashSys=%d\n", m.StackInuse, m.StackSys, m.MSpanInuse, m.MSpanSys, m.MCacheSys, m.BuckHashSys)
    fmt.Printf("NextGC=%d, LastGC=%d, PauseTotalNs=%d, NumGC=%d, EnableGC=%v, DebugGC=%v\n", m.NextGC, m.LastGC, m.PauseTotalNs, m.NumGC, m.EnableGC, m.DebugGC)
    if cstats {
        for i, b := range m.BySize {
            if b.Mallocs == 0 {
                continue
            }
            fmt.Printf("BySize[%d]: Size=%d, Malloc=%d, Frees=%d\n", i, b.Size, b.Mallocs, b.Frees)
        }
    }
    if gc {
        for i := range m.PauseNs {
            fmt.Printf("PauseNs: ")
            fmt.Printf("%d, ", m.PauseNs[(int(m.NumGC)+255+i)%256])
            fmt.Printf("\n")
        }
    }
}

func tdiff(begin, end time.Time) time.Duration {
    d := end.Sub(begin)
    return d
}

func ReadSortWrite(r *bufio.Reader, w *bufio.Writer) (rcnt, wcnt int64) {
    var cp = func(s string) {
        if *tf {
            fmt.Printf("%s", s)
        }
    }
    data := make(StringSlice, 0)
    cp(fmt.Sprintf("read:  "))
    begin := time.Now()
    for {
        str, err := r.ReadString(10) // 0x0A separator = newline
        if err == io.EOF {
            break
        } else if err != nil {
            panic("I/O error")
        }
        //fmt.Printf("%q\n", str[:len(str)-2])
        data = append(data, str[:len(str)-2])
        rcnt++
    }
    end := time.Now()
    d := tdiff(begin, end)
    cp(fmt.Sprintf("time=%0.3f secs\n", float64(d)/1e9))

    // dump memory stats before Sort
    if *mf {
        runtime.ReadMemStats(memstats)
        dump_mstats(memstats, false, false)
    }

    cp(fmt.Sprintf("sort:  "))
    begin = time.Now()
    heapsort.Heapsort(data)
    end = time.Now()
    d = tdiff(begin, end)
    cp(fmt.Sprintf("time=%0.3f secs\n", float64(d)/1e9))

    // dump memory stats after Sort
    if *mf {
        runtime.ReadMemStats(memstats)
        dump_mstats(memstats, false, false)
    }

    cp(fmt.Sprintf("write: "))
    begin = time.Now()
    for _, s := range data {
        if _, err := w.WriteString(s+"\r\n"); err != nil {
            panic(err)
        }
        wcnt++
    }
    if err := w.Flush(); err != nil { panic(err) }
    end = time.Now()
    d = tdiff(begin, end)
    cp(fmt.Sprintf("time=%0.3f secs\n", float64(d)/1e9))
    return
}

func main() {
    var paths []string

    flag.ErrHelp = errors.New("flag: help requested")

    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage of %s [-m][-v] unsortedfile sortedfile\n", os.Args[0])
        flag.PrintDefaults()
    }

    flag.Parse()
    if flag.NArg() != 2 {
        fmt.Printf("usage: example [-m][-v] unsortedfile sortedfile\n")
        os.Exit(1)
    }
    for i := 0; i < flag.NArg(); i++ {
        paths = append(paths, flag.Arg(i))
    }
    memstats = new(runtime.MemStats)

	in, err := os.Open(paths[0])
    if err != nil {
        fmt.Println("error opening file ", err)
        os.Exit(1)
    }
    defer in.Close()

    out, err := os.Create(paths[1])
    if err != nil {
        fmt.Println("error opening file ", err)
        os.Exit(1)
    }
    defer out.Close()
    r := bufio.NewReader(in)
    w := bufio.NewWriter(out)

    rcnt, wcnt := ReadSortWrite(r, w)
    if *vf {
        fmt.Printf("rcnt=%d, wcnt=%d\n", rcnt, wcnt)
    }

    if *mf {
        runtime.ReadMemStats(memstats)
        dump_mstats(memstats, false, false)
    }
}
