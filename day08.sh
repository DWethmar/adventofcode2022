#! /bin/bash

CHALLENGE=day08

go build -o $(pwd)/bin/$CHALLENGE ./$CHALLENGE/*.go

echo "Tests"
cat $CHALLENGE/input_test_01.txt | bin/$CHALLENGE

echo "Solution"
cat $CHALLENGE/input.txt | bin/$CHALLENGE
