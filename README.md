HeapSort
========

Introduction
------------

There are three versions of HeapSort in this repository:

1. heapsort.go
2. c/heapsort.c
3. js/heapsort.js (incomplete)


Approach
--------
My approach to the assignment was in 4 phases.

First, I refreshed my memory of HeapSort by looking on the web and in Knuth Volume 3. It's basically a two phase sort the first part being the setup of a heap, a complete binary tree with an associated invariant that orders the nodes. The second part of the algorithm is a selection sort where the highest value (at the root) is moved from the end of a vector and the heap is rebuilt again, with the length being 1 less, to satisfy the the invariant.

Second, I looked at a number of implementations and they were all about the same. Of course some were clearer and better written than others. I coded a simple version in Go to get a feel for the algorithm.

Third, there were some decisions to make about the algorithm. It had to use location 0 of a slice and many of the examples start at location 1. That's an easy adjustment. Another decision to make was to use the Go sort package interface called Interface. Go's libraries have a public interface for sort algorithms called, creatively enough, "Interface". This interface only has 3 functions defined:

	type Interface interface {
		Len() int
		Less(i, j int) bool
		Swap(i, j int)
	}

I decided to support the interface. It caused me to have to change my algorithm a bit because I used copy operations to move nodes in the heap and the interface only has a swap operation. With some regret I changed my version of HeapSort to use swaps instead of copies.

Finally, I wrote the 3 versions submitted.

The Go version of HeapSort
--------------------------
The Go version is easily built and tested by issuing the commands:

	go build
	go test

The C version of HeapSort
-------------------------
The C version is easily built and tested by issuing the commands:

	make
	./test

A Note on the JavaScript Version of HeapSort
--------------------------------------------

I had originally wanted to do a JavaScript version in addition to the Go version. I've done some JavaScript programming and enjoy it. I knew the biggest issues would be the file I/O but I knew there was a new [File][1] API  that Google was working on and this would give me a chance to explore that. I also figured that if that didn't work out I could find some kind of command line/CLI way to ay least read "stdin" and write "stdout" and I could write a simple shell script to drive it.

Writing the js version of HeapSort was easy. It didn't take too long to figure out the new native file I/O support and get that turning over for file reads. Writes were harder and not documented at the website that I used to get up to speed on the native file API. I then looked at another web page on the same site about the [Filesystem][2] APIs. I was surprised to find out that:

**"In April 2014, it was announced on public-webapps that the Filesystem API spec should be considered dead. Other browsers have showed little interest in implementing it."**

*I decided not to complete the .js version*

The JavaScript sort has been tested and the "index.html" web page allows the user to upload a file using the [File][1] API by dragging it onto the web page. The file is sorted using my JavaScrip HeapSort and the first 10 lines of the sorted file are displayed on the web page after the sort finishes.

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
