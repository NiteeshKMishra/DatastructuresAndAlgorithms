package algorithms

func SelectionSort(list []int) {
	for i := 0; i < len(list); i++ {
		minIndex := i
		for j := i + 1; j < len(list); j++ {
			if list[j] < list[minIndex] {
				minIndex = j
			}
		}
		if i != minIndex {
			list[i], list[minIndex] = list[minIndex], list[i]
		}
	}
}
