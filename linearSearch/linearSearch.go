package linearSearch

func LinearSearch(array []int, value int) int {
	for i := 0; i < len(array); i++ {
		if array[i] == value {
			return i
		}
	}
	return -1
}
