package main

import (
	"io"
	"strings"
	"testing"
)

const testInput = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

const testInput2 = `U 1
U 1
U 1
U 1`

const testInput3 = `U 4`

const testInput4 = `U 1
R 1
U 1
R 1
U 1
R 1
`

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

func TestMoveRope(t *testing.T) {
	type args struct {
		input io.Reader
		debug bool
	}
	tests := []struct {
		name                             string
		args                             args
		wantUniquePositionsVisitedByTail int
		wantErr                          bool
	}{
		{
			name: "example",
			args: args{
				input: strings.NewReader(testInput),
				debug: true,
			},
			wantUniquePositionsVisitedByTail: 13,
			wantErr:                          false,
		},
		// {
		// 	name: "example2",
		// 	args: args{
		// 		input: strings.NewReader(testInput2),
		// 		debug: false,
		// 	},
		// 	wantUniquePositionsVisitedByTail: 3,
		// 	wantErr:                          false,
		// },
		// {
		// 	name: "example3",
		// 	args: args{
		// 		input: strings.NewReader(testInput3),
		// 		debug: false,
		// 	},
		// 	wantUniquePositionsVisitedByTail: 3,
		// 	wantErr:                          false,
		// },
		// {
		// 	name: "example4",
		// 	args: args{
		// 		input: strings.NewReader(strings.Join([]string{"U 1", "R 1", "U 1", "R 1", "U 1", "R 1"}, "\n")),
		// 		debug: true,
		// 	},
		// 	wantUniquePositionsVisitedByTail: 3,
		// 	wantErr:                          false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUniquePositionsVisitedByTail, err := MoveRope(tt.args.input, tt.args.debug)
			if (err != nil) != tt.wantErr {
				t.Errorf("MoveRope() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotUniquePositionsVisitedByTail != tt.wantUniquePositionsVisitedByTail {
				t.Errorf("MoveRope() = %v, want %v", gotUniquePositionsVisitedByTail, tt.wantUniquePositionsVisitedByTail)
			}
		})
	}
}

func TestFollow(t *testing.T) {
	type args struct {
		point  *Point
		target *Point
	}
	tests := []struct {
		name  string
		args  args
		wantX int
		wantY int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := Follow(tt.args.point, tt.args.target)
			if gotX != tt.wantX {
				t.Errorf("Follow() gotX = %v, want %v", gotX, tt.wantX)
			}
			if gotY != tt.wantY {
				t.Errorf("Follow() gotY = %v, want %v", gotY, tt.wantY)
			}
		})
	}
}
