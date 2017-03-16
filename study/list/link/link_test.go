package link

import (
	"fmt"
	"testing"
)

func TestLink(t *testing.T) {
	l := NewLink()
	fmt.Println("新建一个链表,链表长度为：", l.Length())
	for i := 0; i < 20; i++ {
		err := l.Insert(i, i)
		if err != nil {
			fmt.Println("插入值错误！", err)
		}
	}
	fmt.Println("插入值！链表长度为：", l.Length())
	l.Remove(10)
	fmt.Println("删除第10个节点！链表长度为：", l.Length())
	fmt.Println("查找第5个节点！", l.FindK(5))
	fmt.Println("查找值为10的节点！", l.FindVal(10))
	fmt.Println("查找值为9的节点！", l.FindVal(9))
}
