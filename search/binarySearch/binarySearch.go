package search

func BinarySearch(array []int, value int, low int, high int) int {
	middle := (high + low) / 2

	if high < low || len(array) == 0 {
		return -1
	} else if array[middle] > value {
		return BinarySearch(array, value, low, middle-1)
	} else if array[middle] < value {
		return BinarySearch(array, value, middle+1, high)
	}

	return middle
}
