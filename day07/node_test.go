package main

import (
	"reflect"
	"testing"
)

func TestNode_Child(t *testing.T) {
	type fields struct {
		Parent   *Node
		Children []*Node
		Name     string
		Size     int
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Node
	}{
		{
			name: "test",
			fields: fields{
				Parent: nil,
				Children: []*Node{
					{
						Parent:   nil,
						Children: []*Node{},
						Name:     "test",
						Size:     12,
					},
				},
				Name: "x",
				Size: 0,
			},
			args: args{
				name: "test",
			},
			want: &Node{
				Parent:   nil,
				Children: []*Node{},
				Name:     "test",
				Size:     12,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				Parent:   tt.fields.Parent,
				Children: tt.fields.Children,
				Name:     tt.fields.Name,
				Size:     tt.fields.Size,
			}
			if got := n.Child(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.Child() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNode(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "test",
			args: args{
				name: "test",
			},
			want: &Node{
				Parent:   nil,
				Children: []*Node{},
				Name:     "test",
				Size:     0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNode(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodeSize(t *testing.T) {
	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				node: exampleTree,
			},
			want: 48381165,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NodeSize(tt.args.node); got != tt.want {
				t.Errorf("NodeSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIterateNodes(t *testing.T) {
	type args struct {
		node *Node
		f    func(node *Node) bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			IterateNodes(tt.args.node, tt.args.f)
		})
	}
}
