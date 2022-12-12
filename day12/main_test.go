package main

import (
	"reflect"
	"testing"
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

func TestPoint_Distance(t *testing.T) {
	type fields struct {
		X int
		Y int
	}
	type args struct {
		p2 *Point
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Point{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := p.Distance(tt.args.p2); got != tt.want {
				t.Errorf("Point.Distance() = %v, want %v", got, tt.want)
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

func TestGetPathToSignal(t *testing.T) {
	type args struct {
		grid Grid
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				grid: Grid{
					[]rune("Sabqponm"),
					[]rune("abcryxxl"),
					[]rune("accszExk"),
					[]rune("acctuvwj"),
					[]rune("abdefghi"),
				},
			},
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPathToSignal(tt.args.grid); got != tt.want {
				t.Errorf("GetPathToSignal() got = %v, want %v", got, tt.want)
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
