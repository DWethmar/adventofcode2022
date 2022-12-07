package main

import (
	"io"
	"reflect"
	"testing"
)

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

func TestCreateTree(t *testing.T) {
	type args struct {
		in io.Reader
	}
	tests := []struct {
		name     string
		args     args
		wantNode *Node
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNode, err := CreateTree(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNode, tt.wantNode) {
				t.Errorf("CreateTree() = %v, want %v", gotNode, tt.wantNode)
			}
		})
	}
}

func TestSumNodesWithSizeLT(t *testing.T) {
	type args struct {
		node *Node
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
			if got := SumNodesWithSizeLT(tt.args.node); got != tt.want {
				t.Errorf("SumNodesWithSizeLT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSmallestFolderThatWouldFreeUpSpace(t *testing.T) {
	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{
				node: &Node{
					Name: "",
					Children: []*Node{
						{
							Name: "a",
							Children: []*Node{
								{
									Name: "e",
									Children: []*Node{
										{
											Name: "i",
											Size: 584,
										},
									},
								},
								{
									Name: "f",
									Size: 29116,
								},
								{
									Name: "g",
									Size: 2557,
								},
								{
									Name: "h.lst",
									Size: 62596,
								},
							},
						},
						{
							Name: "b.txt",
							Size: 14848514,
						},
						{
							Name: "c.txt",
							Size: 8504156,
						},
						{
							Name: "d",
							Children: []*Node{
								{
									Name: "j",
									Size: 4060174,
								},
								{
									Name: "d.log",
									Size: 8033020,
								},
								{
									Name: "d.ext",
									Size: 5626152,
								},
								{
									Name: "k",
									Size: 7214296,
								},
							},
						},
					},
				},
			},
			want: 24933642,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SmallestFolderThatWouldFreeUpSpace(tt.args.node); got != tt.want {
				t.Errorf("SmallestFolderThatWouldFreeUpSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
