package list

type List struct {
	array []interface{}
	size int
}

func New() *List {
	return &List{}
}

func (list *List) Get(index int) interface{} {
	if index > list.size {
		return false
	}
	return list.array[index]
}

func (list *List) Push(value interface{}) {
	list.size++

	newArray := make([]interface{}, list.size)
	copy(newArray, list.array)

	list.array = newArray
	list.array[list.size - 1] = value
}

func (list *List) Pop() {
	list.size--

	newArray := make([]interface{}, list.size)
	copy(newArray, list.array[:list.size])

	list.array = newArray
}