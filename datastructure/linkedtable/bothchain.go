package linkedtable

type DNode struct {
	data interface{}
	prev *DNode
	next *DNode
}

type DList struct {
	size int
	head *DNode
	tail *DNode
}

func (dList *DList) Init() {
	dList.size = 0
	dList.head = nil
	dList.tail = nil
}

func (dList *DList) Append(node *DNode) {
	if dList.size == 0 {
		dList.head = node
		dList.tail = node

		node.prev = nil
		node.next = nil
	} else {
		node.prev = dList.tail
		node.next = nil

		dList.tail.next = node
		dList.tail = node
	}
	dList.size++
}

func (dList *DList) InsertNext(elmt *DNode, node *DNode) bool {
	if dList.tail == elmt {
		dList.Append(elmt)
	} else {
		node.prev = elmt
		node.next = elmt.next

		elmt.next = node
		node.next.prev = node
	}
	dList.size++
	return true
}

func (dList *DList) insertPrev(elmt *DNode, node *DNode) bool {
	if dList.head == elmt {
		node.next = dList.head
		dList.head.prev = node

		node.prev = nil
		dList.head = node
	} else {
		prev := elmt.prev
		dList.InsertNext(prev, node)
	}
	dList.size++
	return true
}

func (dList *DList) Remove(elmt *DNode) bool {
	prev := elmt.prev
	next := elmt.next

	if dList.head == elmt {
		dList.head = next
	} else {
		prev.next = next
	}

	if dList.tail == elmt {
		dList.tail = prev
	} else {
		next.prev = prev
	}

	dList.size--
	return true
}

func (dList *DList) Search(data interface{}) *DNode {
	node := dList.head
	for i := 1; i < dList.size; i++ {
		if node.data == data {
			return node
		}
		node = node.next
	}
	return nil
}
