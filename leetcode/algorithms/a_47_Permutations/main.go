package main

import (
	"fmt"
)

/*
    生成排列 数组中可能包含重复值。

    字典序生成排列，用的树做启发。同层的相同值的元素，可以消去。构造一棵树后，即可遍历得出 uniqPer
 */
func permuteUnique(nums []int) [][]int {
	return newTree(nums).permList()
}

type node struct {
	value int
	subValueList []int // 另存一份
	children []*node
}

// 递归构造结点。知道最后一个
func newNode(value int, valueList []int) *node {

	// 先构造 childRen
	nodup := delDuplicates(valueList)
	var children []*node

	if len(valueList) >= 1 {
		children = make([]*node, len(nodup))
		for k, v := range nodup { // 构造dup. 去掉v
			children[k] = newNode(v, delElem(valueList, v))
		}
	}

	return &node{
		value,
		nodup,
		children,
	}
}

type tree struct {
	root *node
}

// 从带有重复元素的list中，构造一棵树. 可用来生成排列数
func newTree(valueList []int) *tree {
	return &tree{newNode(0, valueList)}
}

func (t *tree) permlist(node *node, prefix []int) [][]int {
	var (
		plist = [][]int{}
	)


	if node.children == nil { // 终结了
		plist = [][]int{prefix}
	} else {
		for _, child := range node.children {
			newPrefix := make([]int, len(prefix)+1)
			copy(newPrefix, prefix)
			newPrefix[len(prefix)] = child.value
			plist = append(plist, t.permlist(child, newPrefix)...)
		}
	}

	return plist
}

// 遍历出排列数
func (t *tree) permList() [][]int {
	return t.permlist(t.root, []int{})
}

func delDuplicates(valueList []int) []int {

	newList := []int{}

	var i, j int
	// i 位置的元素，与它之前的是否有重复. 如果不重复，计入
	for i=0; i<len(valueList); i++ {
		for j=0; j<i; j++ {
			if valueList[i] == valueList[j] {
				goto outoffirstloop
			}
		}
		newList = append(newList, valueList[i])
outoffirstloop:
	}

	return newList
}

// 在 valueList 里边，删除一个 value.
func delElem(valueList []int, value int) []int {
	var (
		index int
		l = len(valueList)
		newList = make([]int, l-1)
	)

	// 先找到要删除的 index
	for i:=0; i<l; i++ {
		if valueList[i] == value {
			index = i
			break
		}
	}

	if index == l - 1 {
		copy(newList, valueList[:l-1])
	} else {
		tem := make([]int, l)
		copy(tem, valueList)
		copy(newList, append(valueList[0:index], valueList[index+1:]...))
		copy(valueList, tem)
	}

	return newList
}


func main() {
	fmt.Println(delDuplicates([]int{1,2,3,4,2,5,7,8,5}))
	v := []int{1,2,3}
	fmt.Println(delElem(v, 2), v)

	tree := newTree([]int{1,2,3,2})
	fmt.Println(tree.permList())
}
