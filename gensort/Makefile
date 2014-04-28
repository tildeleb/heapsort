all: gensort valsort

gensort: gensort.c rand16.o rand16.h
	gcc -o gensort gensort.c rand16.o -lz

rand16.o: rand16.c rand16.h
	gcc -c rand16.c

valsort: valsort.c rand16.o rand16.h
	gcc -o valsort valsort.c rand16.o -lz

clean:
	-rm *.o
	-rm valsort
	-rm gensort
