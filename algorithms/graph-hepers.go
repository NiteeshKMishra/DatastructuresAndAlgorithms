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
