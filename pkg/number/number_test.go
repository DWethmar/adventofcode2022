package number

import (
	"reflect"
	"testing"
)

func TestGetAllIntsFromString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{s: "1"},
			want: []int{1},
		},
		{
			name: "1 2",
			args: args{s: "1 2"},
			want: []int{1, 2},
		},
		{
			name: "1 2 30",
			args: args{s: "1 2 30"},
			want: []int{1, 2, 30},
		},
		{
			name: "1 2 30 4",
			args: args{s: "1 2 30 4"},
			want: []int{1, 2, 30, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllIntsFromString(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllIntsFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{s: "1"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MustAtoi(tt.args.s); got != tt.want {
				t.Errorf("MustAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
