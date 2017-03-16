package stack

import "errors"

//Element 任意类型，方便更改
type Element int

//Node  堆栈的节点
type Node struct {
	data Element
	next *Node
}

//NewStack 新建一个堆栈
func NewStack() *Node {
	return &Node{}
}

//IsEmpty 返回堆栈是否为空
func (n *Node) IsEmpty() bool {
	return n.next == nil
}

//Push  入栈
func (n *Node) Push(val Element) {
	s := &Node{data: val}
	s.next = n.next
	n.next = s
}

//Pop 出栈
func (n *Node) Pop() (val Element, err error) {
	if n.IsEmpty() == true {
		return val, errors.New("堆栈已空！")
	}
	val = n.next.data
	n.next = n.next.next
	return val, nil
}

//Length 返回堆栈长度
func (n *Node) Length() int {
	p := n
	i := 0
	for p.next != nil {
		p = p.next
		i++
	}
	return i
}
