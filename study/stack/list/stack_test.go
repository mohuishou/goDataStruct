package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack()
	fmt.Println("新建一个堆栈！", s)

	val, err := s.Pop()
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < 21; i++ {
		err = s.Push(Element(i))
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("入栈操作:", s)
	val, err = s.Pop()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("出栈：", val, s)
}
