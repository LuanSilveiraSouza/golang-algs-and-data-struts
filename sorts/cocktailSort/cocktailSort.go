package sorts

func CocktailSort(array []int) []int {
	swapped := true
	var i int

	for swapped {
		swapped = false
		i = 0

		for i < len(array)-1 {
			if array[i] > array[i+1] {
				swapped = true
				aux := array[i]
				array[i] = array[i+1]
				array[i+1] = aux
			}

			i++
		}

		for i > 0 {
			if array[i] < array[i-1] {
				swapped = true
				aux := array[i]
				array[i] = array[i-1]
				array[i-1] = aux
			}

			i--
		}
	}

	return array
}
