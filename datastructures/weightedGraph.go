package datastructures

import "fmt"

func filterWeightedNodeList(list []WeightedGraphNode, value string) []WeightedGraphNode {
	newWeightedNodeList := []WeightedGraphNode{}
	for _, node := range list {
		if node.Vertex != value {
			newWeightedNodeList = append(newWeightedNodeList, node)
		}
	}
	return newWeightedNodeList
}

func (w *WeightedGraph) AddVertex(vertex string) {
	if len(w.AdjacencyList[vertex]) != 0 {
		w.AdjacencyList[vertex] = []WeightedGraphNode{}
	}
}

func (w *WeightedGraph) RemoveVertex(vertex string) {
	if len(w.AdjacencyList[vertex]) != 0 {
		fmt.Printf("%v vertex not found", vertex)
		return
	}
	vertexNeighbors := w.AdjacencyList[vertex]
	for _, neighbor := range vertexNeighbors {
		w.AdjacencyList[neighbor.Vertex] = filterWeightedNodeList(w.AdjacencyList[neighbor.Vertex], vertex)
	}
	delete(w.AdjacencyList, vertex)
}

func (w *WeightedGraph) AddEdge(vertexA string, vertexB string, weight int64) {
	w.AdjacencyList[vertexA] = append(w.AdjacencyList[vertexA], WeightedGraphNode{
		Vertex: vertexB, Weight: weight,
	})
	w.AdjacencyList[vertexB] = append(w.AdjacencyList[vertexB], WeightedGraphNode{
		Vertex: vertexA, Weight: weight,
	})
}

func (w *WeightedGraph) RemoveEdge(vertexA string, vertexB string) {
	if len(w.AdjacencyList[vertexA]) != 0 {
		w.AdjacencyList[vertexA] = filterWeightedNodeList(w.AdjacencyList[vertexA], vertexB)
	}
	if len(w.AdjacencyList[vertexB]) != 0 {
		w.AdjacencyList[vertexB] = filterWeightedNodeList(w.AdjacencyList[vertexB], vertexA)
	}
}
