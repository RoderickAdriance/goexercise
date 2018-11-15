package datastructure

type CNnode struct {
	data interface{}
	next *CNnode
}

type CList struct {
	size int
	head *CNnode
}

func (cList *CList) Init() {
	cList.size = 0
	cList.head = nil
}

func (cList *CList) Append(data interface{}) bool {
	if data == nil {
		return false
	}

	node := new(CNnode)
	node.data = data

	if cList.size == 0 {
		cList.head = node
	} else {
		item := cList.head
		for ; item.next != item; item = item.next {
		}
		item.next = node //放到尾部
	}

	node.next = cList.head //尾部挂到头部
	cList.size++
	return true
}

func (cList *CList) InsertNext(elmt *CNnode, data interface{}) bool {
	if elmt == nil {
		return false
	}

	node := new(CNnode)

	node.next = elmt.next //elmt后面车厢，挂在新车厢后面
	elmt.next = node      // elmt后面挂上新车厢

	cList.size++
	return true
}

func (cList *CList) Remove(elmt *CNnode) bool {
	if elmt == nil {
		return false
	}

	item := cList.head
	for ; item.next != elmt; item = item.next {
	} //获得elmt的前一个

	item.next = elmt.next

	cList.size--
	return true
}
