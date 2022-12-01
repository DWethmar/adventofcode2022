package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_getLargestSum(t *testing.T) {
	type args struct {
		reader io.Reader
		topN   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// {
		// 	name: "test 1",
		// 	args: args{
		// 		reader: strings.NewReader(strings.Join([]string{"1", "2", "3", "", "2", "2", "3", "", "1", "2", "3"}, "\n")),
		// 		topN:   1,
		// 	},
		// 	want: []int{7},
		// },
		{
			name: "test 2",
			args: args{
				reader: strings.NewReader(strings.Join([]string{"1", "2", "3", "", "2", "2", "3", "", "100", "2", "3", "", "1"}, "\n")),
				topN:   2,
			},
			want: []int{105, 7},
		},
		// {
		// 	name: "test 3",
		// 	args: args{
		// 		reader: strings.NewReader(""),
		// 		topN:   3,
		// 	},
		// 	want: []int{0, 0, 0},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLargestSums(tt.args.reader, tt.args.topN); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLargestSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
