package datastructure

type Stack struct {
	list *List
}

func (stack *Stack) Init() {
	list := new(List)
	stack.list = list
}

func (stack *Stack) Push(data interface{}) {
	list := stack.list
	list.Insert(data)
	list.Init()
}

func (stack *Stack) Pop() interface{} {
	list := stack.list
	return list.Remove(0)
}
