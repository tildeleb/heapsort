// Copyright © 2014 Lawrence E. Bakst. All rights reserved.
#include <stdlib.h>
#include <stdio.h>
#include "heapsort.h"

void		Pl(Interface *d, char* str);
long		rbetween(long a, long b);
bool		verify(Interface *d);
void		fill(Interface* d, long a , long b);
bool		TestTrivial();
bool		TestExtended();


// Pl prints a linear list of a heap of values and indicies underneath them.
void
Pl(Interface *d, char* str) {
	Int i, n = d->Len(d);

	fprintf(stderr, "%s: ", str);
	for (i = 0; i < n; i++) {
		fprintf(stderr, "%02ld ", (Int)(d->data[i]));
	};
	fprintf(stderr, "\n");
	fprintf(stderr, "%s: ", str);
	for (i = 0; i < n; i++) {
		fprintf(stderr, "%02ld ", i);
	};
	fprintf(stderr, "\n\n");
};

long
rbetween(long a, long b) {
	long r, ret, diff;
	double rf, r2, r3;

        r = random();
        rf = (double)r/(double)2147483647;
        diff = b - a + 1;
        r2 = rf * (double)diff;
        r3 = r2 + (double)a;
        //fmt.Printf("rbetween: a=%d, b=%d, rf=%f, diff=%f, r2=%f, r3=%f\n", a, b, rf, diff, r2, r3)
        ret = (long)r3;
        return ret;
};

// verify that a slice is in order, returns false for error, not in order
bool
verify(Interface* d) {
	Int v, pv = -1;
	Int i;

	//fprintf(stderr, "d->Len(d)=%ld, d->MaxLen(d)=%ld\n", d->Len(d), d->MaxLen(d));
    for (i = 0; i < d->Len(d); i++) {
		v = (Int)d->Get(d, i);
		//fprintf(stderr, "i=%ld, v=%ld ", i, v);
        if (i == 0) {
            pv = v;
            continue;
        }
        if (pv > v) {
            fprintf(stderr, "Sort out of order\n");
            return true;
        };
    };
    return false;
};

// fill the Interface with randome numbers between a and b inclusive. 
void
fill(Interface* d, long a , long b) {
	Int i;

    for (i = 0; i < d->MaxLen(d); i++) {
    	d->Add(d, (void *)rbetween(a, b), 0);
    }
}

bool
TestTrivial() {
    Int s[] = {60, 94, 66, 44, 43, 68, 7, 16, 10, 30, 52, 81, 22, 38, 32};
    //Interface d = {Len, IntLess, Swap, StrAdd, StrGet, 0, 0, (void *)s, sizeof(Int), sizeof(s)/sizeof(Int), sizeof(s)/sizeof(Int)};
    Interface* d = NewIntInterface(sizeof(s)/sizeof(Int));
    Int i;

    for (i = 0; i < sizeof(s)/sizeof(Int); i++) {
    	d->Add(d, (void*)s[i], 0);
    }

    Pl(d, "s");
    HeapSort(d);
    Pl(d, "e");
    return verify(d);
}

bool
TestExtended() {
	long		n;
	Interface*	d;
	Int			i;

    for (i = 0; i < 25; i++) {
        // set to false for fixed pattern or to true for different values each run
        if (true) {
			srandomdev();
        }
        n = rbetween(10, 1000000);
        fprintf(stderr, "%ld: %ld ", i, n);
        d = NewIntInterface(n);
        fill(d, 1, n*3);
        HeapSort(d);
        if (verify(d))
        	return true;
        DelInterface(d);
    }
    fprintf(stderr, "\n");
    return false;
}

int main(int argc, char *argv[]) {
	TestTrivial();
	TestExtended();
}
