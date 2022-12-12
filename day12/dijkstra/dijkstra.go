package dijkstra

func Dijkstra(graph *WeightedGraph, name string) {
	visited := make(map[string]bool)
	heap := &Heap{}

	startNode := graph.GetNode(name)
	startNode.value = 0
	heap.Push(startNode)

	for heap.Size() > 0 {
		current := heap.Pop()
		visited[current.name] = true
		edges := graph.Edges[current.name]
		for _, edge := range edges {
			if !visited[edge.node.name] {
				heap.Push(edge.node)
				if current.value+edge.weight < edge.node.value {
					edge.node.value = current.value + edge.weight
					edge.node.through = current
				}
			}
		}
	}
}
