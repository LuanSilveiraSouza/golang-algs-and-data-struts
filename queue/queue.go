package queue

type Queue struct {
	array []interface{}
	Size int
}

func (queue *Queue) Push(value interface{}) {
	queue.array = append([]interface{}{value}, queue.array...) 

	queue.Size++
}

func (queue *Queue) Pop() interface{} {
	item := queue.array[len(queue.array)-1]
	queue.array = queue.array[:len(queue.array)-2]

	queue.Size--
	return item
}

func (queue *Queue) Front() interface{} {
	return queue.array[len(queue.array)-1]
}