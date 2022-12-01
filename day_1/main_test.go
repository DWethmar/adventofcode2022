package main

import (
	"io"
	"strings"
	"testing"
)

func Test_getLargestSum(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				reader: strings.NewReader("1\n-2\n3\n1\n\n1\n2\n3\n\n1"),
			},
			want: 6,
		},
		{
			name: "test 2",
			args: args{
				reader: strings.NewReader("1\n-2\n3\n1\n\n100"),
			},
			want: 100,
		},
		{
			name: "test 3",
			args: args{
				reader: strings.NewReader("1\n-2\n3\n\n200\n\n1\n"),
			},
			want: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLargestSum(tt.args.reader); got != tt.want {
				t.Errorf("getLargestSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
