package datastructures

func (g *Graph) DFSRecursive(start string) []string {
	var result []string
	var alreadyVisited = make(map[string]bool)
	var dfsRecursive func(node string)

	if len(g.AdjacencyList[start]) != 0 {
		dfsRecursive = func(node string) {
			if alreadyVisited[node] {
				return
			} else {
				result = append(result, node)
				alreadyVisited[node] = true
				nodeNeighbors := g.AdjacencyList[node]
				for _, value := range nodeNeighbors {
					dfsRecursive(value)
				}
			}
		}
		dfsRecursive(start)
	}

	return result
}

func (g *Graph) DFSIterative(start string) []string {
	var result []string
	var alreadyVisited = make(map[string]bool)
	var currentStack []string

	if len(g.AdjacencyList[start]) != 0 {
		currentStack = append(currentStack, start)
		for len(currentStack) != 0 {
			currentNode := currentStack[len(currentStack)-1]
			currentStack = currentStack[:len(currentStack)-1]
			if !alreadyVisited[currentNode] {
				alreadyVisited[currentNode] = true
				result = append(result, currentNode)
				nodeNeighbors := g.AdjacencyList[currentNode]
				currentStack = append(currentStack, nodeNeighbors...)
			}
		}
	}

	return result
}

func (g *Graph) BFS(start string) []string {
	var result []string
	var alreadyVisited = make(map[string]bool)
	var currentQueue []string

	if len(g.AdjacencyList[start]) != 0 {
		currentQueue = append(currentQueue, start)
		for len(currentQueue) != 0 {
			currentNode := currentQueue[0]
			currentQueue = currentQueue[1:]
			if !alreadyVisited[currentNode] {
				alreadyVisited[currentNode] = true
				result = append(result, currentNode)
				nodeNeighbors := g.AdjacencyList[currentNode]
				currentQueue = append(currentQueue, nodeNeighbors...)
			}
		}
	}

	return result
}
