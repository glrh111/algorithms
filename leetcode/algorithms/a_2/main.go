package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func NewListNodeFromList(lst []int) (l *ListNode) {

	l = &ListNode{}
	cl := l

	for index, value := range lst {
		cl.Val = value
		if index != len(lst)-1 {
			cl.Next = &ListNode{}
			cl = cl.Next
		}
	}

	return
}

func (l *ListNode) ToString() (s string) {
	s = ""
	for i:=l; i!=nil; i=i.Next {
		s += fmt.Sprintf("%d ", i.Val)
	}
	s += "\n"
	return
}

func nowValue(l *ListNode) (value int, next *ListNode) {
	if l == nil {
		value = 0
	} else {
		value = l.Val
		next = l.Next
	}
	return
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	retList := &ListNode{}
	currentRetList := retList
	overFlowValue := 0
	for {
		var v1, v2 int
		v1, l1 = nowValue(l1)
		v2, l2 = nowValue(l2)

		thisValue := overFlowValue + v1 + v2
		overFlowValue = thisValue / 10
		thisValue = thisValue % 10

		currentRetList.Val = thisValue

		if l1 == nil && l2 == nil {
			break
		}

		currentRetList.Next = &ListNode{}
		currentRetList = currentRetList.Next
	}

	if overFlowValue > 0 {
		currentRetList.Next = &ListNode{}
		currentRetList = currentRetList.Next
		currentRetList.Val = overFlowValue
	}

	return retList
}

func main() {
	l1 := NewListNodeFromList([]int{2,4,3})
	l2 := NewListNodeFromList([]int{5,6,4})

	fmt.Println(l1.ToString(), l2.ToString())

	fmt.Println(addTwoNumbers(l1, l2).ToString())

}
