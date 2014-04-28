CC=gcc
all:	hsort heapsort_test

heapsort_test:	heapsort.h heapsort.o heapsort_test.o etime.o
	ld -o heapsort_test heapsort.o heapsort_test.o etime.o -lc

hsort:	heapsort.h heapsort.o hsort.o etime.o
	ld -o hsort heapsort.o hsort.o etime.o -lc

etime.o:	etime.h

clean:
	-rm *.o
	-rm hsort
	-rm heapsort_test
