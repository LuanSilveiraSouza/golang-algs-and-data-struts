package lists

type List struct {
	array []interface{}
	Size  int
}

func (list *List) Get(index int) interface{} {
	if index > list.Size {
		return false
	}
	return list.array[index]
}

func (list *List) Push(value interface{}) {
	list.Size++

	newArray := make([]interface{}, list.Size)
	copy(newArray, list.array)

	list.array = newArray
	list.array[list.Size-1] = value
}

func (list *List) Pop() {
	list.Size--

	newArray := make([]interface{}, list.Size)
	copy(newArray, list.array[:list.Size])

	list.array = newArray
}
