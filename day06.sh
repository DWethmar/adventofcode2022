#! /bin/bash

go build -o $(pwd)/bin/day06 ./day06/*.go

echo "Tests - part 1"
# cat day06/input_test_00.txt | bin/day06
cat day06/input_test_01.txt | bin/day06
cat day06/input_test_02.txt | bin/day06
cat day06/input_test_03.txt | bin/day06
cat day06/input_test_04.txt | bin/day06

echo "Tests - part 2"
cat day06/input_test_10.txt | bin/day06
cat day06/input_test_11.txt | bin/day06
cat day06/input_test_12.txt | bin/day06
cat day06/input_test_13.txt | bin/day06
cat day06/input_test_14.txt | bin/day06

echo "Solution"
cat day06/input.txt | bin/day06
