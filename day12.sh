#! /bin/bash

CHALLENGE=day12

go build -o $(pwd)/bin/$CHALLENGE ./$CHALLENGE/*.go

echo "Tests 1" 
cat $CHALLENGE/input_test_01.txt | bin/$CHALLENGE

echo "Tests 2"
cat $CHALLENGE/input.txt | bin/$CHALLENGE