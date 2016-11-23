/*
给定一棵各节点编号为整数1,2,3...n的树(例如,无环连通图），其Prufer编码（Prufer code，不知道有没有标准的译法，
用金山词霸没有查到，用Google也没有搜索到）构造方法如下：
从树中去掉编号值最小的叶子节点（仅与一条边邻接的节点），以及与它邻接的边后，记下与它邻接的节点的编号。
在树中重复这个过程，直到只剩下一个节点（总是编号为n的节点）为止。
记下的n-1个编号序列就是树的Prufer编码。
你的任务是计算给定的树的Prufer编码。

树用下列语法表示：
  T-->(N S)
  S--> T S | empty
  N-->number
就是说，树由一对括号包围一个表示根节点的数，以及其后跟随的任意多个（可能没有）由单个空格分隔的子树组成。
(这里的表示方法类似编译原理中的产生式）根据这个定义，树的根节点也可能是一个叶子节点。


输入：
   由多个测试案例组成。每个案例用输入文件的一行以上述格式表示一棵树。EOF表示输入结束。可以假定1 <= n <= 50。

输出：
   对于每个测试案例产生包含给定树的Prufer编码的单行输出。数之间由单个空格分隔。不得在行末尾添加任何空格。

示例输入：
(2 (6 (7)) (3) (5 (1) (4)) (8))
(1 (2 (3)))
(6 (1 (4)) (2 (3) (5)))

示例输出：
5 2 5 2 6 2 8
2 3
2 1 6 2 6
*/
package main

import (
	"container/heap"
	"container/list"
	"fmt"
)

type (
	node struct {
		val      int
		parent   *node
		children *list.List
	}

	tree struct {
		count int
		root  *node
		//-----
		cur      *node
		building *node
		leafs    leaflist
	}

	leaflist []*node
)

func (this *leaflist) Len() int {
	return len(*this)
}

func (this *leaflist) Swap(x, y int) {
	(*this)[x], (*this)[y] = (*this)[y], (*this)[x]
}

func (this *leaflist) Less(x, y int) bool {
	return (*this)[x].val < (*this)[y].val
}

func (this *leaflist) Push(val interface{}) {
	*this = append(*this, val.(*node))
}

func (this *leaflist) Pop() interface{} {
	cnt := len(*this)
	val := (*this)[cnt-1]
	*this = (*this)[:cnt-1]
	return val
}

func (this *tree) build(line string) {

	var val int
	var num bool

	this.root = nil
	this.cur = nil
	this.building = nil
	this.leafs = nil
	this.count = 0

	putcur := func() {
		if num {
			this.put("", val)
			this.count++
			val = 0
			num = false
		}
	}

	for _, char := range line {
		switch char {
		case '(':
			putcur()
			this.put("(", 0)
		case ')':
			putcur()
			this.put(")", 0)
		default:
			if '0' <= char && char < '9' {
				val = 10*val + int(char-'0')
				num = true
			} else {
				putcur()
			}
		}
	}
}

func (this *tree) put(code string, val int) {

	switch code {
	case "(":
		this.building = &node{parent: this.cur}
		if this.cur == nil {
			this.root = this.building
			this.cur = this.building
		} else {
			if this.cur.children == nil {
				this.cur.children = list.New()
			}
			this.cur.children.PushBack(this.building)
			this.cur = this.building
		}
	case ")":
		if this.cur.children == nil {
			heap.Push(&this.leafs, this.cur) // 没有子节点的是叶子节点
		} else if this.cur.parent == nil && this.cur.children.Len() == 1 {
			heap.Push(&this.leafs, this.cur) // 根节点只有一个子节点时也是叶子节点
		}
		this.cur = this.cur.parent
	default:
		this.building.val = val
	}
}

func (this *tree) prufer() (code []int) {

	for this.count > 1 {

		min := this.leafs[0]
		parent := min.parent

		heap.Remove(&this.leafs, 0)
		this.count--

		if parent != nil {
			code = append(code, parent.val) // 通常情况: 没有子节点的叶子节点
		} else {
			// 特殊情况: 只有一个子节点的根节点作为叶子节点
			cur := min.children.Front()
			code = append(code, cur.Value.(*node).val)
		}

		if parent != nil {
			for cur := parent.children.Front(); cur != nil; cur = cur.Next() {
				if cur.Value.(*node) == min {
					parent.children.Remove(cur)
					if parent.children.Len() == 0 { // 没有子节点的是叶子节点
						parent.children = nil
						heap.Push(&this.leafs, parent)
					} else if parent == this.root && parent.children.Len() == 1 {
						heap.Push(&this.leafs, parent) // 根节点只有一个子节点时也是叶子节点
					}
					break
				}
			}
		}
	}

	return
}

func main() {

	demo := &tree{}

	demo.build("(2 (6 (7)) (3)  (5 (1) (4)) (8))")
	code := demo.prufer()
	fmt.Println(code)
}
