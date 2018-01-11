package main

import "fmt"

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

/*
    対链表里边的每两个节点，交换位置.
    偶数 even  奇数  odd . 偶数，奇数
 */
func swapPairs(head *ListNode) *ListNode {
	var (
		lastodd *ListNode
		lasteven *ListNode
		index = 0
		current = head
	)

	for {
		if current == nil {
			break
		}

		iseven := index % 2 == 0 // 是否是偶数

		//nownode := current
		nextnode := current.Next // 只是一个指针

		if iseven {
			lasteven = current
		} else {  // 奇数。执行位置交换

			lasteven.Next = nextnode
			current.Next = lasteven

			if lastodd != nil {
				lastodd.Next = current

			} else {
				head = current
			}

			lastodd= lasteven

		}

		index++
		current = nextnode
	}
	return head
}

/*
    lasthead -> (knodes) -> nextnode
    反转每 k 个元素
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	var (
		lasthead, nextnode *ListNode
		current = head
		index int // 0
		//nthk int  // 第几対 index / k
		knodes []*ListNode // k 个元素的。
	)

	for {

		if current == nil {
			break
		}

		//nthk = index / k

		nextnode = current.Next

		knodes = append(knodes, current)

		if len(knodes) >= k { // 集齐了k个

			// 交换 knodes 的内部位置
			for i:=k-1; i>=1; i-- {
				knodes[i].Next = knodes[i-1] // 让他们指向上一个
			}
			knodes[0].Next = nextnode

			if lasthead != nil {
				lasthead.Next = knodes[k-1]
			} else {
				head = knodes[k-1]
			}
			lasthead = knodes[0]
			knodes = []*ListNode{}
		}
		index++
		current = nextnode
	}

	return head
}

func main() {
	lst1 := newListNode([]int{1,2,3,4,5})
	lst2 := newListNode([]int{1,3,4,5,6})
	lst1.Print()
	lst2.Print()

	//swapPairs(lst1).Print()
	reverseKGroup(lst1, 3).Print()

}
