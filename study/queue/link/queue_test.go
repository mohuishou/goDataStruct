package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()
	fmt.Println("新建一个队列！", q, "队列长度：", q.Length())
	for i := 1; i < 20; i++ {
		q.Add(Element(i))
	}
	fmt.Println("入队！", q, "队列长度：", q.Length())
	v, e := q.Delete()
	fmt.Println("出队！", v, q, e, "队列长度：", q.Length())
}
