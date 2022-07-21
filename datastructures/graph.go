package datastructures

func filterList(list []string, value string) []string {
	newList := []string{}
	for _, listValue := range list {
		if listValue != value {
			newList = append(newList, listValue)
		}
	}
	return newList
}

func (g *Graph) AddVertex(vertex string) {
	if len(g.AdjacencyList[vertex]) == 0 { //Go returns zero-th value in this case empty array if key not found
		g.AdjacencyList[vertex] = []string{}
	}
}

func (g *Graph) RemoveVertex(vertex string) {
	if len(g.AdjacencyList[vertex]) != 0 {
		vertexEdges := g.AdjacencyList[vertex]
		for _, edge := range vertexEdges {
			g.AdjacencyList[edge] = filterList(g.AdjacencyList[edge], vertex)
		}
		delete(g.AdjacencyList, vertex)
	}
}

func (g *Graph) AddEdge(vertexA string, vertexB string) {
	if len(g.AdjacencyList[vertexA]) == 0 {
		g.AdjacencyList[vertexA] = []string{vertexB}
	} else {
		g.AdjacencyList[vertexA] = append(g.AdjacencyList[vertexA], vertexB)
	}
	if len(g.AdjacencyList[vertexB]) == 0 {
		g.AdjacencyList[vertexB] = []string{vertexA}
	} else {
		g.AdjacencyList[vertexB] = append(g.AdjacencyList[vertexB], vertexA)
	}
}

func (g *Graph) RemoveEdge(vertexA string, vertexB string) {
	if len(g.AdjacencyList[vertexA]) != 0 {
		g.AdjacencyList[vertexA] = filterList(g.AdjacencyList[vertexA], vertexB)
	}
	if len(g.AdjacencyList[vertexB]) != 0 {
		g.AdjacencyList[vertexB] = filterList(g.AdjacencyList[vertexB], vertexA)
	}
}
