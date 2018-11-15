package datastructure

type Set struct {
	list *List
}

func (set *Set) Init() {
	list := new(List)
	set.list = list
	list.Init()
}

func (set *Set) Insert(data interface{}) bool {
	if set.list.IsMember(data) {
		return false
	}
	set.list.Append(data)

	return true
}

func (set *Set) Remove(data interface{}) bool {
	return set.list.RemoveData(data)
}
