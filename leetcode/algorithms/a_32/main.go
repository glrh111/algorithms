package main

import "fmt"

/*
   先实现一个 stack
 */
type node struct {
	value int
	next *node
}

type stack struct {
	first *node
	size int
}

func newStack() *stack {
	return &stack{nil, 0}
}

func (s *stack) pop() (value int, ok bool) {
	if s.size != 0 {
		ok = true
		value = s.first.value
		s.first = s.first.next
		s.size--
	}
	return
}

func (s *stack) push(value int) {
	newNode := &node{value, s.first}
	s.first = newNode
	s.size++
}

// 返回stk顶部的值，但是不 POP
func (s *stack) peek() (value int, ok bool) {
	if s.size != 0 {
		ok = true
		value = s.first.value
	}
	return
}

func (s *stack) isEmpty() bool {
	return s.size == 0
}

type validParr struct {
	index int     // 最长串开始的索引
	longestl int  // 多长？
}

/*
   再实现功能.
   找出输入字符串中，最长的valid括号
   FIXME 有待修改。结果不正确。没能处理诸如 "()(()" 的情况. 所以要做第二遍筛查。

   参考了这里边的实现：
   https://leetcode.com/articles/longest-valid-parentheses/
 */
func longestValidParentheses(s string) int {
	var (
		longestl, longestindex, nowl int
		stk = newStack()
	)

	// push -1
	stk.push(-1)

	for i:=0; i<len(s); i++ {

		switch char := s[i]; char {
		case '(': // push -> stack
			stk.push(i)
		case ')': // 弹出一个
			stk.pop()
			if stk.isEmpty() { // push i
				stk.push(i)
			} else {           // POP 一个出来
				j, _ := stk.peek() // 然后比较大小

				nowl = i - j   // 最长的
				if nowl > longestl {
					longestl = nowl
					longestindex = i
				}
			}
		default:
			panic("")
		}
	}
	longestindex -= longestl - 1
	fmt.Println(longestindex, longestl)
	return longestl
}

func main() {
	longestValidParentheses("()(()()")
}
