package linkedtable

type Object interface{}

type Node struct {
	data Object
	next *Node
}

type List struct {
	size int
	head *Node
	tail *Node
}

func (list *List) Init() {
	list.size = 0
	list.head = nil
	list.tail = nil
}

func (list *List) Append(node *Node) bool {
	if node == nil {
		return false
	}
	//改进版写法
	if list.size == 0 {
		list.head = node
	} else {
		list.tail.next = node
	}
	list.tail = node
	list.size++

	return true
	//旧写法
	/*if list.size == 0{// 无元素的时候添加
		// 这是单链表的第一个元素，也是链表的头部
		list.head=node
		// 同时是单链表的尾部
		list.tail=node
		// 单链表有了第一个元素
		list.size=1
	}else{  // 有元素了再添加
		oldTail := list.tail // node放到尾部元素后面
		oldTail.next=node

		list.tail = node	// node成为新的尾部
		list.size++			// 元素数量增加
	}*/
}

func (list *List) Insert(node *Node) bool {
	if node == nil {
		return false
	}
	node.next = list.head
	list.head = node
	list.size++
	return true
}

func (list *List) InsertPosition(i int, node *Node) bool {
	if node == nil || i > list.size || list.size == 0 {
		return false
	}

	preItem := list.head
	for j := 1; j < i; j++ {
		preItem = preItem.next
	}
	//放到它后面
	node.next = preItem.next
	preItem.next = node

	list.size++
	return true
}

func (list *List) Remove(i int, node *Node) bool {
	if i >= list.size {
		return false
	}
	if i == 0 {
		node := list.head
		list.head = node.next
		// 如果只有一个元素，那尾部也要调整
		if list.size == 1 {
			list.tail = nil
		}
	} else {
		preItem := list.head
		for j := 1; j < i; j++ {
			preItem = preItem.next
		}

		node = preItem.next
		preItem.next = node.next

		if i == (list.size - 1) { // 若删除的尾部，尾部指针需要调整
			list.tail = preItem
		}
	}

	list.size--
	return true
}

func (list *List) Get(i int) *Node {
	if i >= list.size || i < 0 {
		return nil
	}
	item := list.head
	for j := 0; j < i; j++ {
		item = item.next
	}
	return item
}
