Heapsort Project
================

	Lawrence E. Bakst | leb@me.com | +1-408-930-3801

Introduction
------------

There are three versions of Heapsort in this repository:

1. heapsort.go
2. c/heapsort.c
3. js/heapsort.js (incomplete, but hopefully some extra credit for this)


Approach
--------
My approach to the assignment was to proceed in 5 phases.

First, I refreshed my memory of Heapsort by looking on the web and in Knuth Volume 3. It's basically a two phase sort the first part being the setup of a heap, a complete binary tree with an associated invariant that orders the nodes. The second part of the algorithm is a selection sort where the highest value (at the root) is moved from the end of a vector and the heap is rebuilt again, with the length being 1 less, to satisfy the the invariant.

Second, I looked at a number of implementations and they were all about the same. Of course some were clearer and better written than others. I coded a simple version in Go to get a feel for the algorithm.

Third, there were some decisions to make about the algorithm. It had to use location 0 of a slice and many of the examples start at location 1. That's an easy adjustment. Another decision to make was to use the Go sort package interface called Interface. Go's libraries have a public interface for sort algorithms called, creatively enough, "Interface". This interface only has 3 functions defined:

	type Interface interface {
		Len() int
		Less(i, j int) bool
		Swap(i, j int)
	}

I decided to support the interface. It caused me to have to change my algorithm a bit because I used copy operations to move nodes in the heap and the interface only has a swap operation. With some regret I changed my version of Heapsort to use swaps instead of copies.

Forth, I code the 2 versions of Heapsort.

Finally I saved time at the end of the project to explore performance optimizations but all that time was used up coding the C version.

Building the Project
--------------------
The project was coded and tested on a late model Apple MacBook Pro with Quad Core 2.5 GHz processor that has 16 GB of memory and a 512 MB SSD disk. It was running OS X 10.8.5. Everything should build on linux and I usually run and test on Ubuntu 12.04 LTS either running as a VM on my laptop or running on AWS. However, I didn't have time to build and test on linux so it's entirely possible that this project will neither build or run in linux without some tweaks.

The first step is to clone to project from GitHub:
	
	go get github.com/tildeleb/Heapsort

The project can be built by running the following shell script at top level:

	./build.sh

The project can be built by hand as follows:

	# build gensort
	cd gensort
	make
	# generate large datasets
	cd ../dataset
	./gendata.sh
	cd ..

The Go version of Heapsort
--------------------------
The Go version is easily built and tested by issuing the commands:

	go build
	go test

The C version of Heapsort
-------------------------
The C version is easily built and tested by issuing the commands:

	cd c
	make
	./heapsort_test

The JavaScript version of Heapsort (extra credit)
-------------------------------------------------
This JavaScript doesn't need to be built. The file "index.html" in the js directory needs to be opened in a web browser (Chrome and Safari tested). On the Mac you can do the following:

	cd js
	open index.html

After this drag a file from the datasets directory onto the web page and it will be loaded and sorted. The first 10 lines of the sorted file will be displayed on the web page. I could probably make the file writing work on Chrome but it didn't seem to make sense for a "dead" API. See below for more info.

Validating the Go and C Versions
--------------------------------
There is shell script at top level which will run both the Go and C version against 2 datasets, a 16MB dataset and a 1.2GB dataset.

	./validate.sh

Raw Performance Numbers
-----------------------

	validate Go Heapsort with 16MB dataset
	read:  time=0.081 secs
	sort:  time=0.221 secs
	write: time=0.119 secs
	Records: 163840
	Checksum: 1400fed7bb2c2
	Duplicate keys: 0
	SUCCESS - all records are in order
	
	validate Go Heapsort with 1.2GB dataset
	read:  time=7.001 secs
	sort:  time=49.504 secs
	write: time=17.283 secs
	Records: 13107200
	Checksum: 63f9af71da7682
	Duplicate keys: 0
	SUCCESS - all records are in order
	
	validate C Heapsort with 16MB dataset
	read: 0.2463 seconds
	sort: 0.1581 seconds
	write: 0.0928 seconds
	Records: 163840
	Checksum: 1400fed7bb2c2
	Duplicate keys: 0
	SUCCESS - all records are in order
	
	validate C Heapsort with 1.2GB dataset
	read: 2.3454 seconds
	sort: 40.6514 seconds
	write: 14.5061 seconds
	Records: 13107200
	Checksum: 63f9af71da7682
	Duplicate keys: 0
	SUCCESS - all records are in order

Performance Analysis
--------------------
When sorting the 1.2GB dataset the Go version of Heapsort is about 10 seconds slower than the C version. I think this is to be expected. A slice is accessed and because Go is a safe language each slice access must have it's bounds checked. The interface (indirect) calling sequence is also probably slower than the the function pointers that are used in the C version.

I collected some Go memory heap statistics before and after the sort call in the Go version. The raw numbers are:

	Alloc=1982301720, TotalAlloc=4139607384, Sys=2761696040, Lookups=121, Mallocs=26214676, Frees=12386707
	HeapAlloc=1982301720, HeapSys=2538340352, HeapIdle=278167552, HeapInuse=2260172800, HeapReleased=0, HeapObjects=13827969
	StackInuse=90112, StackSys=262144, MSpanInuse=46380880, MSpanSys=46546944, MCacheSys=32768, BuckHashSys=5161664
	NextGC=3641785680, LastGC=1398679266361912170, PauseTotalNs=1608564143, NumGC=24, EnableGC=true, DebugGC=false
	
	sort:  time=49.286 secs
	
	Alloc=1982302112, TotalAlloc=4139607776, Sys=2761696040, Lookups=121, Mallocs=26214684, Frees=12386707
	HeapAlloc=1982302112, HeapSys=2538340352, HeapIdle=278167552, HeapInuse=2260172800, HeapReleased=0, HeapObjects=13827977
	StackInuse=90112, StackSys=262144, MSpanInuse=46380776, MSpanSys=46546944, MCacheSys=32768, BuckHashSys=5161664
	NextGC=3641785680, LastGC=1398679266361912170, PauseTotalNs=1608564143, NumGC=24, EnableGC=true, DebugGC=false

As expected, the statistics confirm that there where almost no allocations during the Heapsort and no garbage collections where performed.

The Go I/O performance is also slower but this could probably be optimized to get closer to the C performance.

Performance Optimizations
-------------------------
Optimizing the performance of sorting algorithms is a known problem. The keys are:

1. Minimize key compares
2. Avoid function call overhead if possible
3. Hit processor caches if at all possible

Minimizing Key Compares
-----------------------
For Heapsort there is technique first discovered by Floyd and mentioned in Kunth to minimize key compares by moving one of the compares out of the main loop in the siftup. I would have liked to explore this but I ran out of time to do so. Also, the technique requires set/get access to the vector of pointers and the Go Interface I used only provides SWAP and not a COPY operation. The C version of the Interface I designed has operations for get and set so my plan was to explore this optimization in C first to see if the expected speedup of 15-20% can be realized.

Avoiding Function Call Overhead
-------------------------------
Since "interfaces" are used in both the Go and C versions of the code there are many indirect function calls. If the key compare function call can be inlined this overhead can be eliminated. It would be interesting to see how much overhead is due to function calls. With the C version macros and tweaking gcc could be used to explore various options for inlining.

Hitting Processor Caches
------------------------
The difference between a L1 cache hit and a worst case complete memory miss is at least 1000:1 if not 10,000:1 or higher. Neither the 16 MB or 1.2GB dataset fit in any of the processor caches on my laptop but it might be possible to find a AWS instance with larger caches, maybe large enough to hold the entire 16 MB dataset. This would be interesting to explore.

In addition all of the record data are accessed indirectly via pointers in a vector. In theory the data could be stored in a single flat vector and offsets used to avoid the indirection. This *might* improve memory performance but it's not clear how much the indirection hurts and how much the indirect pointers pollute the various caches.

Opportunities for Concurrency
-----------------------------
For very large datasets, larger than the 1.2 GB dataset used here it's possible the the first phase of algorithm could be modified to distribute the siftup process using channels and go-routines. However there there are always (significant) synchronization and rendezvous overheads to deal with.

Conclusion
----------
There are significant opportunities performance optimizations to explore but I didn't have time to work on any of them. I had planned to spend a day on this at the end of project but having to write Javascript and C versions of the code took most of the spare time I had allocated for optimization.


A Note on the JavaScript Version of Heapsort
--------------------------------------------

I had originally wanted to do a JavaScript version in addition to the Go version. I've done some JavaScript programming and enjoy it. I'd like to focus on Go and JavaScript programming so it makes sense to do a JavaScript version. 

I knew the biggest issues would be the file I/O to the local filesystem. However I knew there was a new [File][1] API that Google was working on and this would give me a chance to explore that. I knew it worked in Safari and Chrome. I also assumed that if that didn't work out I could find some kind of command line/CLI way to ay least read "stdin" and write "stdout" and then write a simple shell script to drive it.

Writing the js version of Heapsort was easy. It didn't take too long to understand the new native file I/O support and get that turning over for file reads. Writes were harder and not documented at the website that I used to get up to speed on the native file API. I then looked at another web page on the same site about the [Filesystem][2] APIs. I was surprised to find out that:

**"In April 2014, it was announced on public-webapps that the Filesystem API spec should be considered dead. Other browsers have showed little interest in implementing it.**

**Because of this and the lack of local filesystem write support in Safari I decided not to complete the Javascript version. I probably could have completed a Chrome only Javascript version. However dead is dead, and I didn't want to submit a project that used a dead API.**

The JavaScript sort has been tested and the "index.html" web page allows the user to upload a file using the [File][1] API by dragging it onto the web page. The file is sorted using my JavaScrip Heapsort and the first 10 lines of the sorted file are displayed on the web page after the sort finishes.

JavaScript I/O Still Disappoints
--------------------------------
It does not seem easy to find a way to read and write lines from stdin/stdout in JavaScript. Node.js has something but I am not sure it really allows you to read stdin a line at a time.  

I considered setting up a sever on an AWS micro instance I have running all the time and writing a small Go program to host the POST and GET verbs needed for uploading and downloading of files. My worry was this would not be a self contained submission, so I decided against that and to include what I did in JavaScript as "extra credit".

A Note on the [File][1] APIs
----------------------------
I didn't find these APIs too hard to use or understand. In addition to the noted website I also read the specs at the W3C and Kronos web sites.

I am not sure the chained I/O approach to doing file I/O is the best way to approach an I/O API but there are a number of complex issues when doing I/O in JavaScript that is part of a web page experience because of the single threaded nature of most JavaScript and keeping other the user interface running concurrently with performing I/O. Perhaps the upcoming ECMAScript 6 Specification with Promises will allow the API to be revived in a more multithreaded way and without having to chain I/Os during callbacks.

[1]: http://www.html5rocks.com/en/tutorials/file/dndfiles/

[2]: http://www.html5rocks.com/en/tutorials/file/filesystem/#toc-filesystemurls

*Copyright © 2014 Lawrence E. Bakst. All rights reserved.*
