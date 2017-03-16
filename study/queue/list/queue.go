//Package queue 队列的顺序表实现
package queue

import "errors"

//MAXSIZE 储存数据的最大值
const MAXSIZE = 7

//Element 定义数据类型
type Element int

//Node 队列节点
type Node struct {
	data  [MAXSIZE]Element
	rear  int //列头元素位置
	front int //列尾元素位置
}

//NewQueue 新建一个队列
func NewQueue() *Node {
	return &Node{}
}

//Add 入队
func (n *Node) Add(val Element) error {
	if (n.rear+1)%MAXSIZE == n.front {
		return errors.New("队列已满！")
	}
	n.rear = (n.rear + 1) % MAXSIZE
	n.data[n.rear] = val
	return nil
}

//Delete 出队
func (n *Node) Delete() (val Element, err error) {
	if n.front == n.rear {
		return val, errors.New("队空！")
	}
	n.front = (n.front + 1) % MAXSIZE
	return n.data[n.front], nil
}
