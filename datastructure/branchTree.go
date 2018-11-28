package datastructure

import "container/list"

type Tree struct {
	data  interface{}
	left  *Tree
	right *Tree
}

func createTree(data interface{}) *Tree {
	return &Tree{data: data}
}

//先序遍历
func (bt *Tree) PreOrder() []interface{} {
	t := bt
	stack := list.New()
	res := make([]interface{}, 0)

	for t != nil || stack.Len() == 0 {
		//将左边全部遍历,但是右边还没有遍历,所以先放到List中等待右遍历
		for t != nil {
			res = append(res, t.data)
			stack.PushBack(t) //尾部插入
			t = t.left
		}

		if stack.Len() != 0 {
			v := stack.Back() //获取尾部
			t := v.Value.(*Tree)
			t = t.right
			stack.Remove(v) //删除左右全部遍历的树
		}
	}
	return res
}

//中序遍历,   如果元素没有左树就取出它的值,再处理元素的右树
func (bt *Tree) MiddleOrder() []interface{} {
	t := bt
	stack := list.New()
	res := make([]interface{}, 0)

	for t != nil || stack.Len() != 0 {
		for t != nil {
			stack.PushBack(t)
			t = t.left
		}

		if stack.Len() != 0 {
			v := stack.Back()
			t = v.Value.(*Tree)
			res = append(res, t.data)
			t = t.right
			stack.Remove(v)
		}
	}
	return res

}
