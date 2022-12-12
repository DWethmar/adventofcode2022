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
		name     string
		args     args
		wantPath []*Point
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
			wantPath: []*Point{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPath := GetPathToSignal(tt.args.grid); !reflect.DeepEqual(gotPath, tt.wantPath) {
				t.Errorf("GetPathToSignal() = %v, want %v", gotPath, tt.wantPath)
			}
		})
	}
}

func TestWalk(t *testing.T) {
	type args struct {
		grid   Grid
		start  *Point
		end    *Point
		walked []string
	}
	tests := []struct {
		name     string
		args     args
		wantPath []*Point
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPath := Walk(tt.args.grid, tt.args.start, tt.args.end, tt.args.walked); !reflect.DeepEqual(gotPath, tt.wantPath) {
				t.Errorf("Walk() = %v, want %v", gotPath, tt.wantPath)
			}
		})
	}
}
