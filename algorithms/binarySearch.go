package algorithms

import "math"

func BinarySearch(list []int, value int) int {
	if len(list) == 0 {
		return -1
	}
	index := -1
	leftIndex := 0
	rightIndex := len(list)
	for rightIndex >= leftIndex {
		pivot := int(math.Floor(float64(leftIndex+rightIndex) / 2))
		if value == list[pivot] {
			index = pivot
			break
		} else if value < list[pivot] {
			rightIndex = pivot - 1
		} else {
			leftIndex = pivot + 1
		}
	}
	return index
}
