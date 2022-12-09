package main

import (
	"io"
	"reflect"
	"testing"
)

var testGrid = Grid{
	{3, 0, 3, 7, 3},
	{2, 5, 5, 1, 2},
	{6, 5, 3, 3, 2},
	{3, 3, 5, 4, 9},
	{3, 5, 3, 9, 0},
}

func TestGrid_Get(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		g    Grid
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.Get(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Grid.Get() = %v, want %v", got, tt.want)
			}
		})
	}
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

func TestCreateGrid(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name     string
		args     args
		wantGrid Grid
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGrid, err := CreateGrid(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateGrid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotGrid, tt.wantGrid) {
				t.Errorf("CreateGrid() = %v, want %v", gotGrid, tt.wantGrid)
			}
		})
	}
}

func TestCountVisibleTrees(t *testing.T) {
	type args struct {
		grid Grid
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
		wantErr   bool
	}{
		{
			name: "example",
			args: args{
				grid: testGrid,
			},
			wantCount: 21,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCount, err := CountVisibleTrees(tt.args.grid)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountVisibleTrees() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCount != tt.wantCount {
				t.Errorf("CountVisibleTrees() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func TestIsVisibleTree(t *testing.T) {
	type args struct {
		grid Grid
		x    int
		y    int
	}
	tests := []struct {
		name        string
		args        args
		wantVisible bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotVisible := IsVisibleTree(tt.args.grid, tt.args.x, tt.args.y); gotVisible != tt.wantVisible {
				t.Errorf("IsVisibleTree() = %v, want %v", gotVisible, tt.wantVisible)
			}
		})
	}
}

func TestWalkGrid(t *testing.T) {
	type args struct {
		grid Grid
		x    int
		y    int
		dX   int
		dY   int
		f    func(x, y int) Step
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WalkGrid(tt.args.grid, tt.args.x, tt.args.y, tt.args.dX, tt.args.dY, tt.args.f); (err != nil) != tt.wantErr {
				t.Errorf("WalkGrid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHighestScenicScore(t *testing.T) {
	type args struct {
		grid Grid
	}
	tests := []struct {
		name      string
		args      args
		wantScore int
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotScore, err := HighestScenicScore(tt.args.grid)
			if (err != nil) != tt.wantErr {
				t.Errorf("HighestScenicScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotScore != tt.wantScore {
				t.Errorf("HighestScenicScore() = %v, want %v", gotScore, tt.wantScore)
			}
		})
	}
}

func TestScenicScore(t *testing.T) {
	type args struct {
		grid Grid
		x    int
		y    int
	}
	tests := []struct {
		name      string
		args      args
		wantScore int
	}{
		{
			name: "example 1",
			args: args{
				grid: testGrid,
				x:    2,
				y:    1,
			},
			wantScore: 4,
		},
		{
			name: "example 2",
			args: args{
				grid: testGrid,
				x:    2,
				y:    3,
			},
			wantScore: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotScore := ScenicScore(tt.args.grid, tt.args.x, tt.args.y); gotScore != tt.wantScore {
				t.Errorf("ScenicScore() = %v, want %v", gotScore, tt.wantScore)
			}
		})
	}
}
