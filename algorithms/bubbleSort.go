package algorithms

func BubbleSort(list []int) {
	var noSwaps bool
	for i := len(list); i > 0; i-- {
		noSwaps = true
		for j := 0; j < i-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				noSwaps = false
			}
		}
		if noSwaps {
			break
		}
	}
}
