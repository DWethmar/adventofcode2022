package main

var exampleTree = &Node{
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
}
