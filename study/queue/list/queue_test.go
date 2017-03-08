package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()
	fmt.Println("新建一个队列！", q)
	for i := 1; i < 20; i++ {
		err := q.Add(Element(i))
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	fmt.Println("入队！", q)
	v, e := q.Delete()
	fmt.Println("出队！", v, q, e)
}
