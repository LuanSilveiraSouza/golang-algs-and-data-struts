package insertionSort

func Sort(array []int) []int {
	for currentIndex := 1; currentIndex < len(array); currentIndex++ {
		currentValue := array[currentIndex]
		i := currentIndex
		
		for ; i > 0 && array[i-1] >= currentValue; i-- {
			array[i] = 	array[i - 1]
		} 
		array[i] = currentValue
	}

	return array
}