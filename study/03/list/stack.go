package stack

import "errors"

//Element 任意类型，方便更改
type Element int

//MAXSIZE 数组的最大长度
const MAXSIZE = 20

//Node  堆栈的节点
type Node struct {
	data [MAXSIZE]Element
	top  int
}

//NewStack 新建一个堆栈
func NewStack() *Node {
	return &Node{top: -1}
}

//IsFull 返回堆栈是否已满
func (n *Node) IsFull() bool {
	if n.top == MAXSIZE-1 {
		return true
	}
	return false
}

//IsEmpty 返回堆栈是否为空
func (n *Node) IsEmpty() bool {
	if n.top == -1 {
		return true
	}
	return false
}

//Push  入栈
func (n *Node) Push(val Element) error {
	if n.IsFull() == true {
		return errors.New("堆栈已满！")
	}
	n.top++
	n.data[n.top] = val
	return nil
}

//Pop 出栈
func (n *Node) Pop() (val Element, err error) {
	if n.IsEmpty() == true {
		return val, errors.New("堆栈已空！")
	}
	val = n.data[n.top]
	n.top--
	return val, nil
}
