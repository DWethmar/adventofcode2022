package dijkstra

import "fmt"

type node struct {
	Name  string
	Value interface{}
}

type edge struct {
	node   *node
	weight int
}

type Graph struct {
	Nodes []*node
	Edges map[string][]*edge
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: []*node{},
		Edges: make(map[string][]*edge),
	}
}

func (g *Graph) AddNode(name string, value interface{}) {
	if g.getNode(name) != nil {
		panic(fmt.Sprintf("Node %s already exists", name))
	}

	g.Nodes = append(g.Nodes, &node{
		Name:  name,
		Value: value,
	})
}

func (g *Graph) AddEdge(origin, destiny string, weight int) {
	if g.getNode(origin) == nil {
		panic(fmt.Sprintf("origin node %s not found", origin))
	}

	d := g.getNode(destiny)
	if d == nil {
		panic(fmt.Sprintf("destiny node %s not found", destiny))
	}

	g.Edges[origin] = append(g.Edges[origin], &edge{
		d,
		weight,
	})
}

func (g *Graph) getNode(name string) (node *node) {
	for _, n := range g.Nodes {
		if n.Name == name {
			node = n
		}
	}

	return
}

func (g *Graph) Path(origin string, match func(v interface{}) bool) (int, []string) {
	h := newHeap()
	h.push(path{
		value: 0,
		nodes: []string{origin},
	})

	visited := make(map[string]bool)

	for len(*h.values) > 0 {
		// Find the nearest yet to visit node
		p := h.pop()
		name := p.nodes[len(p.nodes)-1]

		node := g.getNode(name)
		if node == nil {
			panic("Node not found")
		}

		if visited[node.Name] {
			continue
		}

		if match(node.Value) {
			return p.value, p.nodes
		}

		for _, e := range g.Edges[node.Name] {
			if !visited[e.node.Name] {
				// We calculate the total spent so far plus the cost and the path of getting here
				h.push(path{
					value: p.value + e.weight,
					nodes: append([]string{}, append(p.nodes, e.node.Name)...),
				})
			}
		}

		visited[node.Name] = true
	}

	return 0, nil
}
