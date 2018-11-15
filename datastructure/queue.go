package datastructure

type Queue struct {
	list *List
}

func (queue *Queue) Init() {
	list := new(List)
	queue.list = list
	list.Init()
}

func (queue *Queue) Enqueue(data interface{}) bool {
	return queue.list.Append(data)
}

func (queue *Queue) Dequeue() interface{} {
	return queue.list.Remove(0)
}
