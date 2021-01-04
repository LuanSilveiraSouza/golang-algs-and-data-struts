package mergeSort

func merge(a []int, b []int) []int {
	array := make([]int, len(a)+len(b))
	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			array[i+j] = a[i]
			i++
		} else {
			array[i+j] = b[j]
			j++
		}
	}

	for i < len(a) {
		array[i+j] = a[i]
		i++
	}

	for j < len(b) {
		array[i+j] = b[j]
		j++
	}

	return array
}

func Sort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	middle := len(array) / 2
	a := Sort(array[:middle])
	b := Sort(array[middle:])

	return merge(a, b)
}
