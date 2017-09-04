package tree

import (
	"fmt"
	"container/list"
)

type Element int

type TreeNode struct {
	Data       Element
	RightChild *TreeNode
	LeftChild  *TreeNode
}

func NewNode(e Element) *TreeNode {
	return &TreeNode{Data: e}
}

func Find(root *TreeNode, data Element) *TreeNode {
	for root != nil {
		if data < root.Data {
			root = root.LeftChild
		} else if data > root.Data {
			root = root.RightChild
		}else {
			return root
		}
	}
	return root

}

func AddNode(root *TreeNode, data Element) *TreeNode {
	if root == nil {
		root = NewNode(data)
	} else {
		if data < root.Data || data == root.Data {
			root.LeftChild = AddNode(root.LeftChild, data)
		}
		if data > root.Data {
			root.RightChild = AddNode(root.RightChild, data)
		}
	}
	return root
}

func FindMin(root *TreeNode) *TreeNode {
	for root != nil && root.LeftChild != nil {
		root = root.LeftChild

	}
	return root
}

func FindMax(root *TreeNode) *TreeNode {
	for root != nil &&root.RightChild != nil {
		root = root.RightChild
	}
	return root
}

func PreOrder(root *TreeNode) {
	if root != nil {
		fmt.Print(root.Data, " ")
		PreOrder(root.LeftChild)
		PreOrder(root.RightChild)
	}
}

func InOrder(root *TreeNode) {
	if root != nil {
		InOrder(root.LeftChild)
		fmt.Print(root.Data, " ")
		InOrder(root.RightChild)
	}
}

func LastOrder(root *TreeNode) {
	if root != nil {
		LastOrder(root.LeftChild)
		LastOrder(root.RightChild)
		fmt.Print(root.Data, " ")
	}
}

func LevelOrder(root *TreeNode) {
	q := list.New()
	q.PushBack(root)
	for q.Len() > 0 {
		l := q.Front()
		r := l.Value
		q.Remove(l)
		if v, ok := r.(*TreeNode); ok && v != nil {
			fmt.Print(v.Data, " ")
			q.PushBack(v.LeftChild)
			q.PushBack(v.RightChild)
		}
	}
}

func Remove(root *TreeNode, data Element) *TreeNode {
	if root == nil {
		fmt.Println("不存在该节点")
	} else {
		if data < root.Data {
			root.LeftChild = Remove(root.LeftChild, data)
		} else if data > root.Data {
			root.RightChild = Remove(root.RightChild, data)
		} else {
			if root.LeftChild != nil && root.RightChild != nil {
				tmp := FindMin(root.RightChild)
				root.Data = tmp.Data
				root.RightChild = Remove(root.RightChild, tmp.Data)
			} else if root.LeftChild == nil {
				root = root.RightChild
			} else {
				root = root.LeftChild
			}
		}
	}
	return root
}
