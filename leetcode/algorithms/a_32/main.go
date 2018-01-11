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

/*
   再实现功能.
   找出输入字符串中，最长的valid括号
   FIXME 有待修改。结果不正确。没能处理诸如 "()(()" 的情况
 */
func longestValidParentheses(s string) int {
	var (
		longestl, longestindex, nowl int
		stk = newStack()
	)
	
	for i:=0; i<len(s); i++ {

		switch char := s[i]; char {
		case '(': // push -> stack
			stk.push(int(char))
		case ')': // 弹出一个
			c, ok := stk.pop()
			if ok && c-int(char)==-1 { // 匹配了一对儿括号
				nowl += 2
				if nowl > longestl {
					longestl = nowl // 这里还得判断一下，TODO 判断括号里边是否有剩余括号。如果有的话，放入缓冲区
					longestindex = i
				}
			} else { // 看看未闭合的括号的数量有没有变化。
				nowl = 0
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
	longestValidParentheses(")()()()(")
}
