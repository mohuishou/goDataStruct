//Package queue 队列的链表实现
package queue

import (
	"errors"
)

//Element 定义数据类型
type Element int

//Node 队列节点
type Node struct {
	data Element
	next *Node //列头元素位置
}

//Length 返回队列长度
func (q Queue) Length() int {
	i := 0
	for q.front != q.rear {
		q.front = q.front.next
		i++
	}
	return i
}

//Queue 队列
type Queue struct {
	rear  *Node
	front *Node
}

//NewQueue 新建一个队列
func NewQueue() *Queue {
	return &Queue{}
}

//Add 入队
func (q *Queue) Add(val Element) {
	n := &Node{data: val}
	if q.front == nil {
		q.front = n
		q.rear = n
		return
	}
	q.rear.next = n
	q.rear = q.rear.next
}

//Delete 出队
func (q *Queue) Delete() (val Element, err error) {
	if q.front == nil {
		return val, errors.New("队列已空")
	}
	n := q.front
	if q.front == q.rear {
		q.front = nil
		q.rear = nil
	} else {
		q.front = q.front.next
	}
	return n.data, nil
}
