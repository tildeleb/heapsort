#!/bin/bash
cd gensort
make
cd ..
cd datasets
./gendata.sh
cd ..
go build -v
cd example
go build -v
cd ../c
make