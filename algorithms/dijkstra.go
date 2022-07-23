package algorithms

import (
	"math"

	"github.com/niteeshkrmishra/dsandalgo/datastructures"
)

func reverseList(list []string) []string {
	newlist := []string{}
	for len(list) > 0 {
		newlist = append(newlist, list[len(list)-1])
		list = list[:len(list)-1]
	}
	return newlist
}

func setupInitialDistance(list map[string][]datastructures.WeightedGraphNode, vertex string) map[string]int64 {
	initialDistances := map[string]int64{}
	for key := range list {
		if key == vertex {
			initialDistances[key] = 0
		} else {
			initialDistances[key] = math.MaxInt64
		}
	}
	return initialDistances
}

func setupPriorityQueue(list map[string]int64, vertex string) datastructures.PriorityQueue {
	priorityQueue := datastructures.PriorityQueue{}
	for value, priority := range list {
		priorityQueue.Enqueue(value, priority)
	}
	return priorityQueue
}

func setupPreviousValues(list map[string]int64) map[string]string {
	previousValues := map[string]string{}
	for key := range list {
		previousValues[key] = ""
	}
	return previousValues
}

func ShortestPath(w *datastructures.WeightedGraph, startVertex string, endVertex string) []string {
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
