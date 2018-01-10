package main

import "fmt"

/*
   ()[] {} 输入可能是左侧的 6个字符，检测括号的位置是否正确。
   用栈来解决

   node1 <- node2 <- node3 <- ... <- node_first
                                     push
                                     pop
                                     first

 */
func isValid(s string) bool {
	var (
		stk = newStack()
		valid = true
		parr = map[string]int{
			"(": 1, ")": 2,
			"[": 3, "]": 4,
			"{": 5, "}": 6,
		}
	)

	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stk.push(int(char))
		case ')', ']', '}':
			c, ok := stk.pop()
			if ! ok || parr[string(char)]-parr[string(c)] != 1 { // 成对的括号，ASCII不是间隔1
				valid = false
				break
			}
		default:
			valid = false
			break
		}
	}
	if stk.size != 0 {
		valid = false
	}
	return valid
}

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

func main() {
	fmt.Println(isValid("({})"))
}


