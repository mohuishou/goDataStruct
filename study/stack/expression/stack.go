package expression

import "errors"

//Element 任意类型，方便更改
type Element struct {
	Expression string
	Lv         int
}

//Node  堆栈的节点
type Node struct {
	Data Element
	Next *Node
}

//NewStack 新建一个堆栈
func NewStack() *Node {
	return &Node{}
}

//IsEmpty 返回堆栈是否为空
func (n *Node) IsEmpty() bool {
	return n.Next == nil
}

//Push  入栈
func (n *Node) Push(val Element) {
	s := &Node{Data: val}
	s.Next = n.Next
	n.Next = s
}

//Pop 出栈
func (n *Node) Pop() (val Element, err error) {
	if n.IsEmpty() == true {
		return val, errors.New("堆栈已空！")
	}
	val = n.Next.Data
	n.Next = n.Next.Next
	return val, nil
}

//Length 返回堆栈长度
func (n *Node) Length() int {
	p := n
	i := 0
	for p.Next != nil {
		p = p.Next
		i++
	}
	return i
}
