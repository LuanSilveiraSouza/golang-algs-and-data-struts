package search

func InterpolationSearch(array []int, value int, low int, high int) int {
	middle := low + ((value - array[low]) * (high - low) / (array[high] - array[low]))

	if array[middle] == value {
		return middle
	} else if array[middle] > value {
		return InterpolationSearch(array, value, low, middle-1)
	} else if array[middle] < value {
		return InterpolationSearch(array, value, middle+1, high)
	}

	return -1
}
