package main

import (
	"io"
	"os"
	"reflect"
	"testing"
)

var testMonkeys = []*Monkey{
	{
		Name:          "Monkey 0",
		StartingItems: []int{79, 98},
		Operation:     "new = old * 19",
		Test:          "divisible by 23",
		TestTrue:      "throw to monkey 2",
		TestFalse:     "throw to monkey 3",
	},
	{
		Name:          "Monkey 1",
		StartingItems: []int{54, 65, 75, 74},
		Operation:     "new = old + 6",
		Test:          "divisible by 19",
		TestTrue:      "throw to monkey 2",
		TestFalse:     "throw to monkey 0",
	},
	{
		Name:          "Monkey 2",
		StartingItems: []int{79, 60, 97},
		Operation:     "new = old * old",
		Test:          "divisible by 13",
		TestTrue:      "throw to monkey 1",
		TestFalse:     "throw to monkey 3",
	},
	{
		Name:          "Monkey 3",
		StartingItems: []int{74},
		Operation:     "new = old + 3",
		Test:          "divisible by 17",
		TestTrue:      "throw to monkey 0",
		TestFalse:     "throw to monkey 1",
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
			got, err := CreateMonkeys(tt.args.input)
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
			PlayKeepAway(1, tt.args.monkeys)
		})
	}
}
