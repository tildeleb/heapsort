CC=gcc
all:	hsort heapsort_test

heapsort_test:	heapsort.o heapsort_test.o
	ld -o heapsort_test heapsort.o heapsort_test.o -lc

hsort:	heapsort.o hsort.o
	ld -o hsort heapsort.o hsort.o -lc
