#! /bin/bash

go build -o $(pwd)/bin/day07 ./day07/*.go

echo "Tests - part 1"
cat day07/input_test_01.txt | bin/day07

echo "solution - part 1"
cat day07/input.txt | bin/day07
