package sorting

import "fmt"

func BubbleSort(array [5]int) [5]int {
	swapped := true

	for swapped {
		swapped = false
		for j := 0; j < 4; j++ {
			if array[j] > array[j + 1] {
				swapped = true
				aux := array[j]
				array[j] = array[j + 1]
				array[j + 1] = aux
			}
		}
		fmt.Println(array)
	}

	return array
}
