package main

import (
	"reflect"
	"testing"

	"github.com/dwethmar/adventofcode2022/day12/dijkstra"
)

func TestPoint_String(t *testing.T) {
	type fields struct {
		X int
		Y int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Point{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := p.String(); got != tt.want {
				t.Errorf("Point.String() = %v, want %v", got, tt.want)
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

func TestPart1(t *testing.T) {
	type args struct {
		grid  Grid
		graph *dijkstra.Graph
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Part1",
			args: args{
				grid: testGrid,
				graph: CreateGraph(testGrid, func(a, b int) bool {
					// the elevation of the destination square can be at most one higher than the elevation of your current square
					return a > b || a == b || a+1 == b
				}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Part1(tt.args.grid, tt.args.graph)
		})
	}
}

func TestMakeGraph(t *testing.T) {
	type args struct {
		grid   Grid
		isEdge func(a, b int) bool
	}
	tests := []struct {
		name string
		args args
		want *dijkstra.Graph
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateGraph(tt.args.grid, tt.args.isEdge); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetHeight(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHeight(tt.args.r); got != tt.want {
				t.Errorf("GetHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNeighbors(t *testing.T) {
	type args struct {
		grid Grid
		p    *Point
	}
	tests := []struct {
		name          string
		args          args
		wantNeighbors []*Point
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNeighbors := GetNeighbors(tt.args.grid, tt.args.p); !reflect.DeepEqual(gotNeighbors, tt.wantNeighbors) {
				t.Errorf("GetNeighbors() = %v, want %v", gotNeighbors, tt.wantNeighbors)
			}
		})
	}
}
