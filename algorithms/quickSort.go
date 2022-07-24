package algorithms

func partition(list []int, startIndex int, endIndex int) int {
	if len(list) < 1 {
		return -1
	}
	pivot := list[startIndex]
	pivotIndex := startIndex
	for i := startIndex + 1; i <= endIndex; i++ {
		if list[i] < pivot {
			pivotIndex++
			list[i], list[pivotIndex] = list[pivotIndex], list[i]
		}
	}
	list[startIndex], list[pivotIndex] = list[pivotIndex], list[startIndex]
	return pivotIndex
}

func quickSort(list []int, startIndex int, endIndex int) []int {
	if startIndex < endIndex {
		pivotIndex := partition(list, startIndex, endIndex)
		quickSort(list, startIndex, pivotIndex-1)
		quickSort(list, pivotIndex+1, endIndex)
	}
	return list
}

func QuickSort(list []int) []int {
	return quickSort(list, 0, len(list)-1)
}
