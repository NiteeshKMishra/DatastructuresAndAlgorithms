package datastructures

import (
	"fmt"
	"math"
)

//Max binary Heap

func bubbleUp(values []int64) []int64 {
	index := len(values) - 1
	parentIndex := int(math.Floor((float64(index) - 1) / 2))
	for parentIndex >= 0 && values[index] > values[parentIndex] {
		values[parentIndex], values[index] = values[index], values[parentIndex]
		index = int(parentIndex)
		parentIndex = int(math.Floor((float64(index) - 1) / 2))
	}
	return values
}

func sinkDown(values []int64) []int64 {
	arrLength := len(values)
	if arrLength > 2 {
		currentIndex := 0
		values[currentIndex], values[arrLength-1] = values[arrLength-1], values[currentIndex]
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
	} else if arrLength == 2 && values[1] > values[0] {
		values[0], values[1] = values[1], values[0]
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
	bh.Values = bh.Values[1:]
	bh.Values = sinkDown(bh.Values)
	bh.Values = bubbleUp(bh.Values)
	return value
}
