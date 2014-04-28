// Copyright © 2014 Lawrence E. Bakst. All rights reserved.
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include "heapsort.h"
#include "etime.h"

void usage(char * name);
bool Sort(FILE* fi, FILE* fo);

bool tflag = false;

bool
Sort(FILE* fi, FILE* fo) {
	Interface* d;
	Int i, len;
	Int recordSize = 100;
	Int maxRecords = 1000000;
	Str s;
	long astart, astop, rstart, rstop, sstart, sstop, wstart, wstop;

	astart = usecs();
	if ( (d = NewStrInterface(maxRecords, recordSize)) == NULL) {
		exit(1);
	}
	astop = usecs();
	//tprint("allocate: ", astart, astop);

	rstart = usecs();
	while (1) {
		if ( (s = NewStr(recordSize)) == NULL)
			exit(1);
		len = fread((void *)s, recordSize, 1, fi);
		if (len <= 0)
			break;
		//fprintf(stdout, "%s", s);
		d->Add(d, s, recordSize);
	}
	rstop = usecs();
	if (tflag)
		tprint("read: ", rstart, rstop);

	//fprintf(stdout, "before sort\n");
	sstart = usecs();
    HeapSort(d);
 	sstop = usecs();
 	if (tflag)
		tprint("sort: ", sstart, sstop);
	//fprintf(stdout, "after sort\n");

	wstart = usecs();
    for (i = 0; i < d->Len(d); i++) {
		if ( (s = d->Get(d, i)) == NULL)
			exit(1);
		//fprintf(stdout, "%s", s);
		len = fwrite((void *)s, recordSize, 1, fo);
		len *= recordSize;
		//fprintf(stderr, "len=%ld, recordSize=%ld\n", len, recordSize);
		if (len != recordSize)
			return false;
	}
 	wstop = usecs();
	if (tflag)
		tprint("write: ", wstart, wstop);
	return true;
}


void
usage(char * name) {
	fprintf(stderr, "%s: [-t] [inputfile|-] [outputfile|-]\n", name);
	exit(1);
}

int
main(int argc, char *argv[]) {
	int i = 0;
	FILE *fin = stdin;
	FILE *fout = stdout;
	char * name = argv[0];

	if (argc > 5) {
		usage(name);
		exit(1);
	}

	if (argc > 1 && argv[1][0] == '-' && argv[1][1] == 't') {
		tflag = true;
		argc--, argv++;
	}
	while (++i < argc) {
		switch (i) {
		case 1:
			if (strcmp("-", argv[i]) == 0)
				continue;
			if ( (fin = fopen(argv[i], "r")) == NULL) {
				fprintf(stderr, "%s: can't open file %s\n", argv[0], argv[i]);
				exit(1);
			}
			break;
		case 2:
			if (strcmp("-", argv[i]) == 0)
				continue;
			if ( (fout = fopen(argv[i], "w")) == NULL) {
				fprintf(stderr, "%s: can't open file %s\n", argv[0], argv[i]);
				exit(1);
			}
			break;
		}
	}
	Sort(fin, fout);
}