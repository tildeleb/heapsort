#!/bin/bash
echo "make gensort"
cd gensort
make
cd ..
echo "generate datasets"
cd datasets
./gendata.sh
cd ..
echo "build Go Heapsort"
go build -v
cd hsort
go build -v
echo "build C Heapsort"
cd ../c
make