package main

import (
	"io"
	"reflect"
	"testing"

	"github.com/dwethmar/adventofcode2022/pkg/iterate"
)

func TestGrid_Get(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		g    Grid
		args args
		want rune
	}{
		{
			name: "Get",
			g:    testGrid,
			args: args{
				x: 0,
				y: 0,
			},
			want: 'S',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.Get(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Grid.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrid_Iterate(t *testing.T) {
	type args struct {
		f func(x, y int, r rune) iterate.Step
	}
	tests := []struct {
		name string
		g    Grid
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.Iterate(tt.args.f)
		})
	}
}

func TestCreateGrid(t *testing.T) {
	type args struct {
		in io.Reader
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
			gotGrid, err := CreateGrid(tt.args.in)
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

func TestFindPoints(t *testing.T) {
	type args struct {
		grid Grid
		p    []rune
	}
	tests := []struct {
		name string
		args args
		want []*Point
	}{
		{
			name: "FindPoints",
			args: args{
				grid: testGrid,
				p:    []rune{'S'},
			},
			want: []*Point{
				{X: 0, Y: 0},
			},
		},
		{
			name: "FindPoints",
			args: args{
				grid: testGrid,
				p:    []rune{'E'},
			},
			want: []*Point{
				{X: 5, Y: 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindPoints(tt.args.grid, tt.args.p...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
