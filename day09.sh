#! /bin/bash

CHALLENGE=day09

go build -o $(pwd)/bin/$CHALLENGE ./$CHALLENGE/*.go

echo "Tests" 
cat $CHALLENGE/input_test_01.txt | bin/$CHALLENGE debug

echo "Solution"
cat $CHALLENGE/input.txt | bin/$CHALLENGE
