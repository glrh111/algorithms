package main

import (
	"time"
	"fmt"
	"math/rand"
)

/*
   生成valid闭合的括号串
   n 表示n対括号
 */
func generateParenthesis(n int) []string {
	return []string{}
}


func rand2(n int) int {
	var (
		seed = time.Now().Nanosecond() // 种子值
		rnd int
	)
	rnd = seed * 1103515245 + 12345
	rnd = (rnd / 65536) % 32768
	return rnd % n
}

// 按照索引，从小到大排序
func mapToSlice(m map[int]int) []int {
	var (
		s = make([]int, len(m))
	)
	for key, value := range m {
		s[key] = value
	}
	return s
}

func main() {
	counter := map[int]int{}

	for i:=0; i<100000; i++ {
		r := rand2(100)
		counter[r] += 1
	}
	s := mapToSlice(counter)
	fmt.Println(s)

	counter2 := map[int]int{}

	for i:=0; i<100000; i++ {
		r := rand.Intn(100)
		counter2[r] += 1
	}
	s2 := mapToSlice(counter2)
	fmt.Println(s2)

}