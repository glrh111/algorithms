package main

import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}


/*
   扩展一下：合并一系列链表。
   改进空间：如果一个列表空了，那么将它移除列表，以便加快速度
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	var (
		currentNode *ListNode = nil
		l *ListNode = nil
		lst = []*ListNode{l1, l2}
		llength = len(lst)
		minIndex = 0
		minValue = 0
		flag = false
	)

	for {
		minIndex = 0
		minValue = 2147483647 + 1
		i := 0
		for {

			if i >= llength || llength == 0 {
				if llength == 0 {
					flag = true
				}
				break
			}

			if lst[i] != nil { // 比较最小值
				if minValue > lst[i].Val {
					minIndex = i
					minValue = lst[i].Val
				}
			} else {          // 移除nil
				if i == llength-1 { // 移除最后一个元素
					lst = lst[0:i]
				} else {
					lst = append(lst[0:i], lst[i+1:]...)
				}
				llength--
				continue
			}
			i++
		}

		if flag {
			break
		}

		// next
		lst[minIndex] = lst[minIndex].Next

		// newCurrentNode
		newNode := &ListNode{minValue, nil}

		if currentNode != nil {
			currentNode.Next = newNode
		} else {
			l = newNode
		}
		currentNode = newNode

	}

	return l
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
	mergeTwoLists(lst1, lst2).Print()
}
