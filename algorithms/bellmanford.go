package algorithms

import (
	"fmt"
	"math"

	"github.com/niteeshkrmishra/dsandalgo/datastructures"
)

func ShortestPathBellmanFord(w *datastructures.WeightedGraph, startVertex string, endVertex string) []string {
	distances := setupInitialDistance(w.AdjacencyList, startVertex)
	previousValues := setupPreviousValues(distances)

	path := []string{}

	for i := 0; i < len(w.AdjacencyList)-1; i++ {
		for vertex, neighbors := range w.AdjacencyList {
			if distances[vertex] != math.MaxInt {
				for _, neighbor := range neighbors {
					neighborVertex := neighbor.Vertex
					distance := neighbor.Weight
					if distances[vertex]+distance < distances[neighborVertex] {
						distances[neighborVertex] = distances[vertex] + distance
						previousValues[neighborVertex] = vertex
					}
				}
			}
		}
	}
	for vertex, neighbors := range w.AdjacencyList {
		if distances[vertex] != math.MaxInt {
			for _, neighbor := range neighbors {
				neighborVertex := neighbor.Vertex
				distance := neighbor.Weight
				if distances[vertex]+distance < distances[neighborVertex] {
					fmt.Println("Graph contains negative cycle")
					return []string{}
				}
			}
		}
	}
	tempVertex := endVertex
	for previousValues[tempVertex] != "" {
		path = append(path, tempVertex)
		tempVertex = previousValues[tempVertex]
	}
	path = append(path, tempVertex)
	return reverseList(path)
}
