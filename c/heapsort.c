// Copyright © 2014 Lawrence E. Bakst. All rights reserved.
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include "heapsort.h"

void
siftup(Interface *d, Int ni, Int ri) {
	Int ci;

	//s := (d).(IntSlice)
	//r := s[ari]

	//    ri := ari
	//fprintf(stderr, "h(s, ni=%ld, ri=%ld), r=%ld\n", ni, ri, (Int)(d->data[ri]));
	while (ri < (ni + 1)/2) {
		ci = ri*2 + 1; // calculate left child node
		//fprintf(stderr, "ci=%ld, ri=%ld, ni=%ld\n", ci, ri, ni);

		if (ci < ni && d->Less(d, ci, ci+1)) { // follow the largest node left or right
			ci++;
			//fprintf(stderr, "ci=%ld\n", ci);
		}
		if (d->Less(d, ci, ri)) // invariant holds // ni < 1 || 
			break;
		// invariant doesn't hold, swap child and root and descent to the next level of tree
		//fprintf(stderr, "descend: Swap(%ld, %ld) = (%ld, %ld)\n", ri, ci, (Int)(d->data[ri]), (Int)(d->data[ci]));
		d->Swap(d, ri, ci);
		ri = ci;
	};
};

// Heapsort sorts data using a heaport algorithm. It implements the standard Go sort.Interface interface.
void
HeapSort(Interface *d) {
	Int i, n = d->Len(d) - 1;
	//fprintf(stderr, "HeapSort: Len=%d\n", d->Len(d));

	// build a heap with max element at index 0 by building a binary tree and repeatedly satisfying the invaraint from the bottom up
	for (i = n / 2; i >= 0; i--) {
		siftup(d, n, i);
		//Pl(d, "i");
	};
	//Pl(d, "m");
	//Pt(d)

	// repeatedly put the root (largest element) at the end of the slice and rebuild the heap with one less element
	while (n > 0) {
		//fmt.Printf("Heapsort: Swap(%d, %d), Swap(%d, %d)\n", 0, n, s[0], s[n])
		d->Swap(d, 0, n--);
		siftup(d, n, 0);
		//Pl(d, "p");
		//Pt(d)
	};
};

Int
Len(Interface *d) {
	return d->len;
};

void
Swap(Interface *d, Int i, Int j) {
	void *t;
	t = d->data[i]; d->data[i] = d->data[j]; d->data[j] = t;
};

void*
Get(Interface *d, Int i) {
	if (i < 0 || i >= d->len)
		return 0;
	return d->data[i];
}

Int
MaxLen(Interface *d) {
	return d->maxLen;
}

bool
IntLess(Interface *d, Int i, Int j) {
	if ((Int)d->data[i] < (Int)d->data[j])
		return true;
	else
		return false;
};

bool
IntAdd(Interface *d, void* p, Int len) {
	//fprintf(stderr, "d->dataLen=%ld, d->len=%ld, d->maxLen=%ld\n", d->dataLen, d->len, d->maxLen);
	d->data[d->len++] = p;
	return false;
}

Interface *
NewIntInterface(Int n) {
	Interface *d;
	if ( (d = (Interface *)calloc(sizeof(Interface), 1)) == NULL) {
		return 0;
	}
	d->Len = Len;
	d->Less = IntLess;
	d->Swap = Swap;
	d->Get = Get;
	d->Add = IntAdd;
	d->MaxLen = MaxLen;
	d->data = (void *)calloc(sizeof(void *), n);
	if (d->data == NULL) {
		return 0;
	};
	d->len = 0;
	d->dataLen = 0; // 0 says data[i] are not pointers to data but the values themselves.
	d->maxLen = n;
	return d;
}

Str
NewStr(Int len) {
	Str s;

	if ( (s = calloc(len, 1)) == NULL) {
		return 0;
	}
	return (Str)s;
}

bool
StrLess(Interface *d, Int i, Int j) {
	if (strncmp((Str)d->data[i], (Str)d->data[j], d->dataLen) < 0)
		return true;
	else
		return false;
};

bool
StrAdd(Interface *d, void* p, Int len) {
	//fprintf(stderr, "d->dataLen=%ld, d->len=%ld, d->maxLen=%ld\n", d->dataLen, d->len, d->maxLen);
	if (d->len >= d->maxLen)
		return true;
	if (len > d->dataLen)
		return true;
	d->data[d->len++] = p;
	return false;
}

Interface *
NewStrInterface(Int n, Int dataLen) {
	Interface *d;
	if ( (d = (Interface *)calloc(sizeof(Interface), 1)) == NULL) {
		return 0;
	}
	d->Len = Len;
	d->Less = StrLess;
	d->Swap = Swap;
	d->Add = StrAdd;
	d->Get = Get;
	d->data = (void *)calloc(sizeof(void *), n);
	if (d->data == NULL) {
		return 0;
	};
	d->len = 0;
	d->dataLen = dataLen;
	d->maxLen = n;
	return d;
}

void
DelInterface(Interface* d) {
	Int i;

	//fprintf(stderr, "d->Len(d)=%ld, d->MaxLen(d)=%ld, dataLen=%ld\n", d->Len(d), d->MaxLen(d), d->dataLen);

	// first if elements are pointers to data, free them
	if (d->dataLen > 0) {
		for (i = 0; i < d->Len(d); i++)
			free(d->Get(d, i));
	}

	// Now free the data array
	free(d->data);

	// Finally free the Interface
	free((void *)d);
}

