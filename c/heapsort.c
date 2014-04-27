#include <stdlib.h>
#include <stdio.h>

typedef char bool;
enum bool {false, true};
typedef long int Int;


typedef struct Interface Interface;
struct Interface {
	int (*Len)(Interface *d);
	bool (*Less)(Interface *d, int i, int j);
	void (*Swap)(Interface *d, int i, int j);
	void* (*data);
	int len;
};

// Pl prints a linear list of a heap of values and indicies underneath them.
void
Pl(Interface *d, char* str) {
	int i, n = d->Len(d);
	fprintf(stderr, "%s: ", str);
	for (i = 0; i < n; i++) {
		fprintf(stderr, "%02ld ", (Int)(d->data[i]));
	};
	fprintf(stderr, "\n");
	fprintf(stderr, "%s: ", str);
	for (i = 0; i < n; i++) {
		fprintf(stderr, "%02d ", i);
	};
	fprintf(stderr, "\n");
};

void
siftup(Interface *d, int ni, int ri) {
	int ci;
	//s := (d).(IntSlice)
	//r := s[ari]

	//    ri := ari
	fprintf(stderr, "h(s, ni=%d, ri=%d), r=%ld\n", ni, ri, (Int)(d->data[ri]));
	while (ri < (ni + 1)/2) {
		ci = ri*2 + 1; // calculate left child node
		fprintf(stderr, "ci=%d, ri=%d, ni=%d\n", ci, ri, ni);

		if (ci < ni && d->Less(d, ci, ci+1)) { // follow the largest node left or right
			ci++;
			fprintf(stderr, "ci=%d\n", ci);
		}
		if (d->Less(d, ci, ri)) // invariant holds // ni < 1 || 
			break;
		// invariant doesn't hold, swap child and root and descent to the next level of tree
		fprintf(stderr, "descend: Swap(%d, %d) = (%ld, %ld)\n", ri, ci, (Int)(d->data[ri]), (Int)(d->data[ci]));
		d->Swap(d, ri, ci);
		ri = ci;
	};
};

// Heapsort sorts data using a heaport algorithm. It implements the standard Go sort.Interface interface.
void
HeapSort(Interface *d) {
	int i, n = d->Len(d) - 1;
	//fprintf(stderr, "HeapSort: Len=%d\n", d->Len(d));

	// build a heap with max element at index 0 by building a binary tree and repeatedly satisfying the invaraint from the bottom up
	for (i = n / 2; i >= 0; i--) {
		siftup(d, n, i);
		Pl(d, "i");
	};
	Pl(d, "m");
	//Pt(d)

	// repeatedly put the root (largest element) at the end of the slice and rebuild the heap with one less element
	while (n > 0) {
		//fmt.Printf("Heapsort: Swap(%d, %d), Swap(%d, %d)\n", 0, n, s[0], s[n])
		d->Swap(d, 0, n--);
		siftup(d, n, 0);
		Pl(d, "p");
		//Pt(d)
	};
};

// verify that a slice is in order
bool
verify(Interface *d) {
    int i, v, pv, n = d->Len(d);
    for (i = 0; i < n; i++) {
        if (i == 0) {
            pv = v;
            continue;
        };
        if (pv > v) {
            fprintf(stderr, "Sort out of order\n");
            return true;
        };
    };
    return false;
};

int
Len(Interface *d) {
	return d->len;
};

void
Swap(Interface *d, int i, int j) {
	void *t;
	t = d->data[i]; d->data[i] = d->data[j]; d->data[j] = t;
};

bool
IntLess(Interface *d, int i, int j) {
	if ((Int)d->data[i] < (Int)d->data[j])
		return true;
	else
		return false;
};

Interface *
NewIntInterface(Int n) {
	Interface *d;
	if ( (d = (Interface *)calloc(sizeof(Interface), 1)) == NULL) {
		return 0;
	}
	d->Len = Len;
	d->Less = IntLess;
	d->Swap = Swap;
	d->data = (void *)calloc(sizeof(void *), n);
	if (d->data == NULL) {
		return 0;
	};
	d->len = n;
	return d;
}

void
TestTrivial() {
    Int s[] = {60, 94, 66, 44, 43, 68, 7, 16, 10, 30, 52, 81, 22, 38, 32};
    Interface d = {Len, IntLess, Swap, (void *)s, sizeof(s)/sizeof(Int)};

    Pl(&d, "s");
    //pt(s)
    HeapSort(&d);
    //s[0] = 100
    //pl(s)
    Pl(&d, "e");
    //Pt(s)
    verify(&d);
}

int main(int argc, char *argv[]) {
	TestTrivial();
}

