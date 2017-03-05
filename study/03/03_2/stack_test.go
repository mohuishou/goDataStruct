package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack()
	fmt.Println("新建一个堆栈！", s, "堆栈长度：", s.Length())

	val, err := s.Pop()
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < 21; i++ {
		s.Push(Element(i))
	}
	fmt.Println("入栈操作:", s, "堆栈长度：", s.Length())
	val, err = s.Pop()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("出栈：", val, s, "堆栈长度：", s.Length())
}
