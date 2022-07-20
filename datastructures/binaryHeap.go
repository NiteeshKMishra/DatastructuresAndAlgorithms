package datastructures

import (
	"fmt"
	"math"
)

//Max binary Heap

func bubbleUp(values []int64) []int64 {
	valueLength := len(values)
	index := valueLength - 1
	for index > 0 {
		parentIndex := int(math.Floor((float64(index) - 1) / 2))
		if values[index] <= values[parentIndex] {
			break
		}
		values[parentIndex], values[index] = values[index], values[parentIndex]
		index = parentIndex
	}
	return values
}

func sinkDown(values []int64) []int64 {
	arrLength := len(values)
	currentIndex := 0
	for {
		currentValue := values[currentIndex]
		leftChildIndex := 2*currentIndex + 1
		rightChildIndex := 2*currentIndex + 2
		tempIndex := currentIndex
		if leftChildIndex < arrLength {
			leftChild := values[leftChildIndex]
			if leftChild > currentValue {
				tempIndex = leftChildIndex
			}
		}
		if rightChildIndex < arrLength {
			leftChild := values[leftChildIndex]
			rightChild := values[rightChildIndex]
			if (tempIndex != currentIndex && rightChild > leftChild) ||
				(tempIndex == currentIndex && rightChild > currentValue) {
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

func (bh *BinaryHeap) MaxBinaryHeapInsert(value int64) {
	bh.Values = append(bh.Values, value)
	bh.Values = bubbleUp(bh.Values)
}

func (bh *BinaryHeap) BinaryHeapExtractMax() int64 {
	if len(bh.Values) == 0 {
		fmt.Println("Empty Heap")
		return -1
	}
	value := bh.Values[0]
	if len(bh.Values) > 1 {
		lastValue := bh.Values[len(bh.Values)-1]
		bh.Values = bh.Values[:len(bh.Values)-1]
		bh.Values[0] = lastValue
		bh.Values = sinkDown(bh.Values)
	} else {
		bh.Values = []int64{}
	}
	return value
}
