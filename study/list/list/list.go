//Package list 线性表
package list

//MAXSIZE 线性表的最大储存限制
const MAXSIZE = 100

//List 线性表节点
type List struct {
	data [MAXSIZE]int
	last int
}

//NewList 返回一个空list
func NewList() List {
	return List{last: -1}
}

//FindK 根据位序K返回相应的元素
func (l List) FindK(k int) int {
	return l.data[k]
}

//FindX 根据元素x返回其第一次出现的位置
func (l List) FindX(x int) int {
	for i := range l.data {
		if l.data[i] == x {
			return i
		}
	}
	return -1
}

//Insert 在位序i前插入一个新的值
func (l List) Insert(x int, k int) List {

	if l.last > MAXSIZE-1 {
		println("表满！")
		return l
	}

	if k < 1 || k > l.last+2 {
		println("插入位置不合法！")
		return l
	}

	for i := l.last; i >= k-1; i++ {
		l.data[i+1] = l.data[i]
	}
	l.data[k-1] = x
	l.last++
	return l
}

//Remove 删除一个值
func (l List) Remove(k int) List {
	if k < 1 || k > l.last+1 {
		println("位置不合法！")
		return l
	}

	for i := k - 1; i < l.last; i++ {
		l.data[i] = l.data[i+1]
	}
	l.last--
	return l
}
