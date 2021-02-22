package stack

type Stack struct {
	array []interface{}
	size int
}

func (stack *Stack) Push(value interface{}) {
	stack.array = append(stack.array, value) 
}

func (stack *Stack) Pop() interface{} {
	item := stack.array[len(stack.array)-1]
	stack.array = stack.array[:len(stack.array)-2]

	return item
}

func (stack *Stack) Top() interface{} {
	return stack.array[len(stack.array)-1]
}