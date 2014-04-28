#!/bin/bash
echo "make gensort"
cd gensort
make
cd ..
echo "generate datasets (may take a minute)"
cd datasets
./gendata.sh
cd ..
echo "build Go Heapsort"
go build -v
go test
echo "test Go Heapsort"
cd hsort
go build -v
echo "build C Heapsort"
cd ../c
make
echo "test C Heapsort"
./heapsort_test
make