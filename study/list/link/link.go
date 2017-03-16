//Package link 链表的实现
package link

import (
	"errors"
)

//Link 链表节点
type Link struct {
	data int
	next *Link
}

//NewLink 初始化一个链表
func NewLink() *Link {
	return &Link{}
}

//Length 返回表长
func (l *Link) Length() int {
	i := 0
	p := l
	for p != nil {
		p = p.next
		i++
	}
	return i
}

//FindK 按照序号查找
func (l *Link) FindK(k int) *Link {
	p := l
	i := 1
	for p != nil && i < k {
		p = p.next
		i++
	}
	if i == k {
		return p
	}
	return nil
}

//FindVal 根据值查找
func (l *Link) FindVal(val int) *Link {
	p := l
	for p != nil && p.data != val {
		p = p.next
	}
	return p
}

//Insert 插入一个值
func (l *Link) Insert(val, key int) error {
	s := &Link{data: val}
	p := l.FindK(key)
	if p == nil {
		return errors.New("key错误！")
	}
	s.next = p.next
	p.next = s
	return nil
}

//Remove 删除一个节点
func (l *Link) Remove(key int) error {
	if key == 1 {
		s := l.next
		l.next = s
		return nil
	}
	p := l.FindK(key - 1)
	if p == nil {
		return errors.New("找不到该节点！")
	}
	s := p.next
	p.next = s.next
	return nil
}
