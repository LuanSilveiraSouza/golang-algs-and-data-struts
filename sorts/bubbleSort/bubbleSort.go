package sorts

func BubbleSort(array []int) []int {
	swapped := true

	for swapped {
		swapped = false
		for j := 0; j < len(array) - 1; j++ {
			if array[j] > array[j + 1] {
				swapped = true
				aux := array[j]
				array[j] = array[j + 1]
				array[j + 1] = aux
			}
		}
	}

	return array
}
