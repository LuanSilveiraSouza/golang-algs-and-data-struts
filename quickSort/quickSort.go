package quickSort

func Sort(array []int, low, high int) []int {
	if (low > high) {
		return array
	}

	pivot := partition(array, low, high)
	Sort(array, low, pivot - 1)
	Sort(array, pivot + 1, high)

	return array
}

func partition(array []int, low, high int) int {
	pivot := array[high]

	for i := low; i < high; i++ {
		if array[i] < pivot {
			array[i], array[low] = array[low], array[i]
			low++
		}
	}

	array[high], array[low] = array[low], array[high]

	return low
}