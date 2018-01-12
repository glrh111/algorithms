package main

import (
	"math/rand"
	"time"
	"fmt"
)

/*
 *   Comparable
 */
// A type, typically a collection, that satisfies sort.Interface can be
// sorted by the routines in this package. The methods require that the
// elements of the collection be enumerated by an integer index.
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
	// 打乱元素
	Shuffle()
	// i < j -1; i == j 0; i > j 1
	Compare(i, j int) int
	Show(lo, hi int)
	// 是否排定顺序
	IsSorted() bool
}

// 整数实现的 Interface
type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p IntSlice) Shuffle() { // 瞎 JB 交换
	l := p.Len()
	rand.Seed(int64(time.Now().Nanosecond()))
	for index, randindex := range rand.Perm(l) {
		if index >= l / 2 { break }
		p.Swap(index, randindex)
	}
}
func (p IntSlice) Compare(i int, j int) (re int) {
	if p[i] > p[j] {
		re = 1
	} else if p[i] == p[j] {
		re = 0
	} else {
		re = -1
	}
	return
}
func (p IntSlice) Show(lo int, hi int) {
	fmt.Println(p[lo:hi+1])
}

func (p IntSlice) IsSorted() (is bool) {
	is = true
	for i:=0; i<p.Len()-1; i++ {
		if p[i] > p[i+1] {
			is = false
			break
		}
	}
	return
}

func generateIntSlice(n int, max int) (p IntSlice) {
	p = make(IntSlice, n)
	rand.Seed(int64(time.Now().Nanosecond()))
	for i:=0; i<n; i++ {
		p[i] = rand.Intn(max + 1)
	}
	p.Shuffle()
	return
}

