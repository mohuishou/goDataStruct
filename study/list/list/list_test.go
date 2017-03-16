package list

import (
	"fmt"
	"testing"
)

var l List

func TestInit(t *testing.T) {
	l = NewList()
	fmt.Println("线性表初始化成功！", l)
}

func TestInsert(t *testing.T) {
	l = l.Insert(10, 2)
	for i := 0; i < 30; i++ {
		l = l.Insert(i, i)
	}
	fmt.Println("线性表插入成功！", l)
}

func TestRemove(t *testing.T) {
	l = l.Remove(10)
	fmt.Println("元素移除成功！", l)
}

func TestFindx(t *testing.T) {
	a := l.FindX(10)
	fmt.Println("元素10查找成功！", a)
}
