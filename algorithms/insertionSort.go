package algorithms

func InsertionSort(list []int) {
	if len(list) > 1 {
		for i := 1; i < len(list); i++ {
			currentValue := list[i]
			insertIndex := i
			for j := i - 1; j >= 0 && list[j] > currentValue; j-- {
				list[j+1] = list[j]
				insertIndex = j
			}
			if insertIndex != i {
				list[insertIndex] = currentValue
			}
		}
	}
}
