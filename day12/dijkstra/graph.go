package dijkstra

type edge struct {
	node   string
	weight int
}

type Graph struct {
	nodes map[string][]edge
}

func NewGraph() *Graph {
	return &Graph{nodes: make(map[string][]edge)}
}

func (g *Graph) AddEdge(origin, destiny string, weight int) {
	g.nodes[origin] = append(g.nodes[origin], edge{node: destiny, weight: weight})
	// g.nodes[destiny] = append(g.nodes[destiny], edge{node: origin, weight: weight})
}

func (g *Graph) edges(node string) []edge {
	return g.nodes[node]
}

func (g *Graph) Path(origin, destiny string) (int, []string) {
	h := newHeap()
	h.push(path{value: 0, nodes: []string{origin}})
	visited := make(map[string]bool)

	for len(*h.values) > 0 {
		// Find the nearest yet to visit node
		p := h.pop()
		node := p.nodes[len(p.nodes)-1]

		if visited[node] {
			continue
		}

		if node == destiny {
			return p.value, p.nodes
		}

		for _, e := range g.edges(node) {
			if !visited[e.node] {
				// We calculate the total spent so far plus the cost and the path of getting here
				h.push(path{value: p.value + e.weight, nodes: append([]string{}, append(p.nodes, e.node)...)})
			}
		}

		visited[node] = true
	}

	return 0, nil
}
