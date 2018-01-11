package main

import "fmt"

/*
   1, 2, 3, 4, 5, 6, 7
   移除 ln 结尾第n个结点

   只有遍历一遍才知道 size，而且只能遍历一遍。那么把 node 转存在slice 里边吗？

 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var (
		nodelst = []*ListNode{}
		size = 0
		currentnode = head
	)
	for {
		if currentnode == nil {
			break
		}

		nodelst = append(nodelst, currentnode)
		currentnode = currentnode.Next
		size++
	}

	if n == size { // 移除头部元素
		head = head.Next // 完成移除
	} else {
		nodelst[size-n-1].Next = nodelst[size-n].Next
	}

	return head
}

type ListNode struct {
	Val int
	Next *ListNode
}

func (l *ListNode) Print() {
	var cl = l
	for {
		if cl == nil {
			break
		}
		fmt.Printf("%d, ", cl.Val)
		cl = cl.Next
	}
	fmt.Println()
}

func newListNode(lst []int) *ListNode {
	var (
		cu *ListNode = nil
		l *ListNode = nil
	)
	for index, value := range lst {
		newNode := &ListNode{value, nil}
		if index == 0 {
			l = newNode
		} else {
			cu.Next = newNode
		}
		cu = newNode
	}
	return l
}

func main() {
	lst1 := newListNode([]int{1,2,3,4,5,7})
	lst2 := newListNode([]int{1,3,4,5,6})
	lst1.Print()
	lst2.Print()
	removeNthFromEnd(lst2, 5).Print()
}
