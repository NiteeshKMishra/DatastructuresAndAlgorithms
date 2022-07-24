package algorithms

import (
	"math"
	"strconv"
)

const NUM_BASE = 10

func getNumAtPosition(number int, position int) int {
	var numAtPosition int
	var tempNum int = number
	for i := 1; i <= position; i++ {
		numAtPosition = tempNum % NUM_BASE
		tempNum = int(math.Floor(float64(tempNum) / NUM_BASE))
	}
	return numAtPosition
}

func flattenList(list [][]int) []int {
	var newList []int
	for _, item := range list {
		newList = append(newList, item...)
	}
	return newList
}

func getMaxNumLength(list []int) int {
	maxNum := list[0]
	for _, item := range list {
		if item > maxNum {
			maxNum = item
		}
	}
	return len(strconv.Itoa(maxNum))
}

func RadixSort(list []int) []int {
	if len(list) <= 1 {
		return list
	}
	sortedList := list
	maxNumLength := getMaxNumLength(list)
	for i := 0; i <= maxNumLength; i++ {
		buffers := make([][]int, NUM_BASE)
		for _, item := range sortedList {
			indexAt := getNumAtPosition(item, i)
			buffers[indexAt] = append(buffers[indexAt], item)
		}
		sortedList = flattenList(buffers)
	}
	return sortedList
}
