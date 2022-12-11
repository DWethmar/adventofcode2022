package main

import (
	"io"
	"math"
	"os"
	"reflect"
	"testing"
)

var testMonkeys = []*Monkey{
	{
		Name:          "Monkey 0",
		Items: []int{79, 98},
		Operation:     "new = old * 19",
		Test:          23,
		TestTrue:      2,
		TestFalse:     3,
	},
	{
		Name:          "Monkey 1",
		Items: []int{54, 65, 75, 74},
		Operation:     "new = old + 6",
		Test:          19,
		TestTrue:      2,
		TestFalse:     0,
	},
	{
		Name:          "Monkey 2",
		Items: []int{79, 60, 97},
		Operation:     "new = old * old",
		Test:          12,
		TestTrue:      1,
		TestFalse:     3,
	},
	{
		Name:          "Monkey 3",
		Items: []int{74},
		Operation:     "new = old + 3",
		Test:          17,
		TestTrue:      0,
		TestFalse:     1,
	},
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func TestCreateMonkeys(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []*Monkey
		wantErr bool
	}{
		{
			name: "example",
			args: args{
				input: func() io.Reader {
					f := "input_test_01.txt"
					r, err := os.Open(f)
					if err != nil {
						t.Fatalf("failed to open %s: %v", f, err)
					}
					return r
				}(),
			},
			want:    testMonkeys,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := CreateMonkeys(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateMonkeys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for i, want := range tt.want {
				if !reflect.DeepEqual(want, got[i]) {
					t.Errorf("CreateMonkeys() = \n %+v, want \n %+v", got[i], want)
				}
			}
		})
	}
}

func TestPlayKeepAway(t *testing.T) {
	type args struct {
		monkeys []*Monkey
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{
				monkeys: testMonkeys,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PlayKeepAway(tt.args.monkeys, func(worryLvl int) int {
				return int(math.Floor(float64(worryLvl) / 3))
			})
		})
	}
}

func Test_part1(t *testing.T) {
	type args struct {
		monkeys []*Monkey
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			part1(tt.args.monkeys)
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		monkeys []*Monkey
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "example",
			args: args{
				monkeys: testMonkeys,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			part2(tt.args.monkeys, 10)
		})
	}
}
