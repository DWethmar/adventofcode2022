#! /bin/bash

echo 'test'
go build -o $(pwd)/bin/day05 ./day05/*.go && bin/day05 $(pwd)/day05/input_test.txt 

echo 'solution'
go build -o $(pwd)/bin/day05 ./day05/*.go && bin/day05 $(pwd)/day05/input.txt 
