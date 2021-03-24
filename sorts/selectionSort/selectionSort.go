package sorts

func SelectionSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		small := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[small] {
				small = j
			}
		}

		if array[i] != array[small] {
			array[i] += array[small]
			array[small] = array[i] - array[small]
			array[i] -= array[small]
		}

	}

	return array
}
