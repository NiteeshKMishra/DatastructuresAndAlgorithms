package algorithms

import (
	"math"

	"github.com/niteeshkrmishra/dsandalgo/datastructures"
)

func ShortestPathDijkstra(w *datastructures.WeightedGraph, startVertex string, endVertex string) []string {
	distances := setupInitialDistance(w.AdjacencyList, startVertex)
	previousValues := setupPreviousValues(distances)
	priorityQueue := setupPriorityQueue(distances, startVertex)
	path := []string{}

	for len(priorityQueue.Values) > 0 {
		currentNode := priorityQueue.Dequeue()
		currentVertex := currentNode.Value
		if currentVertex == endVertex {
			tempVertex := endVertex
			for previousValues[tempVertex] != "" {
				path = append(path, tempVertex)
				tempVertex = previousValues[tempVertex]
			}
			path = append(path, tempVertex)
			break
		}
		if distances[currentVertex] != math.MaxInt {
			for _, neighbor := range w.AdjacencyList[currentVertex] {
				neighborVertex := neighbor.Vertex
				distance := distances[currentVertex] + neighbor.Weight
				if distance < distances[neighborVertex] {
					distances[neighborVertex] = distance
					previousValues[neighborVertex] = currentVertex
					priorityQueue.Enqueue(neighborVertex, distance)
				}
			}
		}
	}
	return reverseList(path)
}
