package algorithms

func LinerSearch(list []int, value int) int {
	index := -1
	for i := 0; i < len(list); i++ {
		if list[i] == value {
			index = i
			break
		}
	}
	return index
}
