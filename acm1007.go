/*
相对顺序错误的项目对的个数是一种衡量序列“无序度”(unsortedness，应该是一个数学术语，在《线性代数》课程中学习过，
但是不记得是不是这么叫)的方法。
比如说，字母序列DAABEC的无序度是5，因为字母D比它右边的4个字母大，字母E比它右边的1个字母大。
这种衡量序列无序度的方法叫做“逆序数”。字母序列AACEDGG只有一个逆序(E和D),它接近有序；
而字母序列ZWQM有6个逆序(它的无序度是最大的－－完全是顺序的相反面).
现在你负责对DNA序列（只含有四个字母A,C,G,T）进行编目。
然而，不是要按照字母顺序进行排序，而是按照“有序度”进行排列，从“最有顺序的”到“最没有顺序的”。
所有的输入串长度相等。

【输入数据】
第一行含有两个整数：正整数n(0 < n <= 50)表示字符串的长度；正整数m(0 < m <= 100)表示给定的字符串的个数。
随后m行都是长度为n的字符串。

【输出数据】
输出输入串列表，从“最有序的”到“最无序的”。两个字符串的有序度一样的时候，按照原顺序输出。（稳定的排序）

【输入样例】
10 6
AACATGAAGG
TTTTGGCCAA
TTTGGCCAAA
GATCAGATTT
CCCGGGGGGA
ATCGATGCAT

【输出样例】
CCCGGGGGGA
AACATGAAGG
GATCAGATTT
ATCGATGCAT
TTTTGGCCAA
TTTGGCCAAA
*/
package main

import (
	"fmt"
	"sort"
)

type (
	item struct {
		dna    string
		unsort int
	}
	itemlist []*item
)

const (
	dna = "TGCA"
)

func (this itemlist) Len() int {
	return len(this)
}

func (this itemlist) Swap(x, y int) {
	this[x], this[y] = this[y], this[x]
}

func (this itemlist) Less(x, y int) bool {
	return this[x].unsort < this[y].unsort
}

func unsortness(str string) (unsort int) {

	count := []int{0, 0, 0, 0}

	for _, char := range str {

		idx := 0
		for rune(dna[idx]) != char {
			idx++
		}

		for i := 0; i < idx; i++ {
			unsort += count[i]
		}

		count[idx]++
	}

	return
}

func demo() {

	data := []string{
		"AACATGAAGG",
		"TTTTGGCCAA",
		"TTTGGCCAAA",
		"GATCAGATTT",
		"CCCGGGGGGA",
		"ATCGATGCAT",
	}

	list := make(itemlist, len(data))
	for idx, str := range data {
		list[idx] = &item{dna: str, unsort: unsortness(str)}
	}

	sort.Sort(list)

	for _, item := range list {
		fmt.Printf("%v: %v\n", item.dna, item.unsort)
	}
}

func main() {
	demo()
}
