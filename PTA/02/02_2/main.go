package main

import "fmt"

func main() {

	a := NewNode()
	b := NewNode()

	read(a)
	read(b)
	res := Multi(a, b)
	printNode(res)
	res = Add(a, b)
	printNode(res)
}

func read(a *Node) {
	var (
		k    int
		cofe int
		exp  int
	)
	fmt.Scan(&k)
	for i := 0; i < k; i++ {
		fmt.Scan(&cofe)
		fmt.Scan(&exp)
		a.Add(cofe, exp)
	}
}

func printNode(a *Node) {
	if a.next == nil {
		fmt.Print(a.coef, " ", a.exp)
	}
	for a.next != nil {
		fmt.Print(a.next.coef, " ", a.next.exp)
		a = a.next
		if a.next != nil {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

//Multi 多项式乘法
func Multi(a, b *Node) *Node {
	res := NewNode()
	tmp := b
	for a.next != nil {
		tmp = b
		for tmp.next != nil {
			cofe := a.next.coef * tmp.next.coef
			exp := a.next.exp + tmp.next.exp
			n := res.FindExp(exp)
			if n != nil {
				cofe = cofe + n.coef
				if cofe == 0 {
					res.remove(exp)
				} else {
					n.coef = cofe
				}
			} else {
				res.Add(cofe, exp)
			}
			tmp = tmp.next
		}
		a = a.next
	}
	return res
}

//Add 多项式加法
func Add(a, b *Node) *Node {
	res := NewNode()
	for a.next != nil && b.next != nil {
		switch {
		case a.next.exp > b.next.exp:
			res.Add(a.next.coef, a.next.exp)
			a = a.next
		case a.next.exp == b.next.exp:
			res.Add(a.next.coef+b.next.coef, a.next.exp)
			a = a.next
			b = b.next
		case a.next.exp < b.next.exp:
			res.Add(b.next.coef, b.next.exp)
			b = b.next
		}
	}

	if a.next != nil {
		for a.next != nil {
			res.Add(a.next.coef, a.next.exp)
			a = a.next
		}
	}

	if b.next != nil {
		for b.next != nil {
			res.Add(b.next.coef, b.next.exp)
			b = b.next
		}
	}

	return res
}

//Node 节点
type Node struct {
	coef int
	exp  int
	next *Node
}

//NewNode 新建节点
func NewNode() *Node {
	return &Node{}
}

func (n *Node) remove(exp int) {
	p := n
	tmp := n
	for p != nil && p.exp != exp {
		tmp = p
		p = p.next
	}
	if p != nil {
		tmp.next = p.next
	}
}

//FindExp 根据指数查找值
func (n Node) FindExp(exp int) *Node {
	p := &n
	p = p.next
	for p != nil && p.exp != exp {
		p = p.next
	}

	return p
}

//Add 有序的添加一个节点
func (n *Node) Add(coef, exp int) {
	if coef == 0 {
		return
	}
	p := n
	t := &Node{coef: coef, exp: exp}
	for p.next != nil && p.next.exp > exp {
		p = p.next
	}
	if p.next != nil {
		t.next = p.next
	}
	p.next = t
}
