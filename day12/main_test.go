package main

import (
	"reflect"
	"testing"

	"github.com/dwethmar/adventofcode2022/day12/dijkstra"
)

var testGrid = Grid{
	[]rune("Sabqponm"),
	[]rune("abcryxxl"),
	[]rune("accszExk"),
	[]rune("acctuvwj"),
	[]rune("abdefghi"),
}

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

func Test_part1(t *testing.T) {
	type args struct {
		grid  Grid
		graph *dijkstra.Graph
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			part1(tt.args.grid, tt.args.graph)
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		grid  Grid
		graph *dijkstra.Graph
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			part2(tt.args.grid, tt.args.graph)
		})
	}
}

func TestMakeGraph(t *testing.T) {
	type args struct {
		grid Grid
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
			if got := MakeGraph(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPathToSignal(t *testing.T) {
	type args struct {
		graph *dijkstra.Graph
		start *Point
		end   *Point
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []string
	}{
		// {
		// 	name: "test",
		// 	args: args{
		// 		graph: MakeGraph(testGrid),
		// 		start: &Point{X: 0, Y: 0},
		// 		end:   &Point{X: 5, Y: 2},
		// 	},
		// 	want:  31,
		// 	want1: []string{},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetPathToSignal(tt.args.graph, tt.args.start, tt.args.end)
			if got != tt.want {
				t.Errorf("GetPathToSignal() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetPathToSignal() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGetHeight(t *testing.T) {
	type args struct {
		grid Grid
		p    *Point
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
			if got := GetHeight(tt.args.grid, tt.args.p); got != tt.want {
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
