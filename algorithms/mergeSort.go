package algorithms

import "math"

func merge(leftList []int, rightList []int) []int {
	resultList := []int{}
	i, j := 0, 0
	for i < len(leftList) && j < len(rightList) {
		if leftList[i] < rightList[j] {
			resultList = append(resultList, leftList[i])
			i++
		} else if rightList[j] < leftList[i] {
			resultList = append(resultList, rightList[j])
			j++
		} else {
			resultList = append(resultList, leftList[i])
			resultList = append(resultList, rightList[j])
			i++
			j++
		}
	}
	for i < len(leftList) {
		resultList = append(resultList, leftList[i])
		i++
	}
	for j < len(rightList) {
		resultList = append(resultList, rightList[j])
		j++
	}
	return resultList
}

func mergeSort(list []int) []int {
	if len(list) <= 1 {
		return list
	}
	pivot := int(math.Floor(float64(len(list)) / 2))
	leftList := mergeSort(list[:pivot])
	rightList := mergeSort(list[pivot:])
	return merge(leftList, rightList)
}

func MergeSort(list []int) []int {
	return mergeSort(list)
}
