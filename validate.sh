#!/bin/bash
# first validate the Go Version
cd example
echo "Start Heapsort validation run"
echo ""
echo "validate Go Heapsort with 16MB dataset"
./example -t ../datasets/l163840 out
../gensort/valsort out
rm out
echo ""
echo "validate Go Heapsort with 1.2GB dataset"
./example -t ../datasets/l12800k out
../gensort/valsort out
rm out
echo ""
cd ..

# next validate the C Version
cd c
echo "validate C Heapsort with 16MB dataset"
hsort -t ../datasets/l163840 out
../gensort/valsort out
rm out
echo ""
echo "validate C Heapsort with 1.2GB dataset"
hsort -t ../datasets/l12800k out
../gensort/valsort out
rm out
cd ..
echo ""
echo "Finish Heapsort validation run"
