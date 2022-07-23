package datastructures

import (
	"fmt"
	"math"
)

//Implemented using MinBinary Heap i.e lowest value has the most priority

func bubbleDown(values []PriorityNode) []PriorityNode { //bubble lowest priority at the top
	valuesLength := len(values)
	currentIndex := valuesLength - 1
	for currentIndex > 0 {
		parentIndex := int(math.Floor((float64(currentIndex) - 1) / 2))
		if values[parentIndex].Priority < values[currentIndex].Priority {
			break
		}
		values[currentIndex], values[parentIndex] = values[parentIndex], values[currentIndex]
		currentIndex = parentIndex
	}
	return values
}

func sinkUp(values []PriorityNode) []PriorityNode { //reaarange elements after removing a element from top
	arrLength := len(values)
	currentIndex := 0
	for {
		currentValue := values[currentIndex]
		leftChildIndex := 2*currentIndex + 1
		rightChildIndex := 2*currentIndex + 2
		tempIndex := currentIndex
		if leftChildIndex < arrLength {
			leftChild := values[leftChildIndex]
			if leftChild.Priority < currentValue.Priority {
				tempIndex = leftChildIndex
			}
		}
		if rightChildIndex < arrLength {
			leftChild := values[leftChildIndex]
			rightChild := values[rightChildIndex]
			if (tempIndex != currentIndex && rightChild.Priority < leftChild.Priority) ||
				(tempIndex == currentIndex && rightChild.Priority < currentValue.Priority) {
				tempIndex = rightChildIndex
			}
		}
		if tempIndex == currentIndex {
			break
		}
		values[currentIndex], values[tempIndex] = values[tempIndex], values[currentIndex]
		currentIndex = tempIndex
	}
	return values
}

func (p *PriorityQueue) Enqueue(value string, priority int64) {
	newNode := PriorityNode{
		Value:    value,
		Priority: priority,
	}
	p.Values = append(p.Values, newNode)
	p.Values = bubbleDown(p.Values)
}

func (p *PriorityQueue) Dequeue() *PriorityNode {
	if len(p.Values) == 0 {
		fmt.Println("Empty priority queue")
		return nil
	}
	value := p.Values[0]
	if len(p.Values) > 1 {
		lastValue := p.Values[len(p.Values)-1]
		p.Values = p.Values[:len(p.Values)-1]
		p.Values[0] = lastValue
		p.Values = sinkUp(p.Values)
	} else {
		p.Values = []PriorityNode{}
	}
	return &value
}
