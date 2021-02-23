package stack

type Stack struct {
	array []interface{}
	Size int
}

func (stack *Stack) Push(value interface{}) {
	stack.array = append(stack.array, value) 
	stack.Size++
}

func (stack *Stack) Pop() interface{} {
	item := stack.array[len(stack.array)-1]
	stack.array = stack.array[:len(stack.array)-2]

	stack.Size--
	return item
}

func (stack *Stack) Top() interface{} {
	return stack.array[len(stack.array)-1]
}