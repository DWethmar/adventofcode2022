#! /bin/bash

echo 'test'
go build -o $(pwd)/bin/day04 ./day04/*.go && bin/day04 $(pwd)/day04/input_test.txt 

echo 'solution'
go build -o $(pwd)/bin/day04 ./day04/*.go && bin/day04 $(pwd)/day04/input.txt 