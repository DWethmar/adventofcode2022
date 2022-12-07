#! /bin/bash

go build -o $(pwd)/bin/day07 ./day07/*.go

echo "Tests"
cat day07/input_test_01.txt | bin/day07

echo "Solution"
cat day07/input.txt | bin/day07
