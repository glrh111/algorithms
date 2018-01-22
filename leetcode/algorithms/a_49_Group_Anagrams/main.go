package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	var (
		// key 是排序后的单词
		uniList = make(map[string][]string)
		re = [][]string{}
	)
	for _, s := range strs {
		key := sortString(s)
		valueList := uniList[key]
		if valueList == nil {
			valueList = []string{}
		}
		valueList = append(valueList, s)
		uniList[key] = valueList
	}
	// 整理 汇总
	for _, value := range uniList {
		re = append(re, value)
	}

	return re
}

/*
   三向切分快速排序
 */
func _partition(r []rune, lo int, hi int) (int, int) {
	var (
		lt, gt, i = lo+1, hi, lo+1
	)
	for {
		if i>gt {
			break
		}
		if r[i] == r[lo] { // 相等
			i++
		} else if r[i] > r[lo] { // 换到后边
			r[gt], r[i] = r[i], r[gt]
			gt--
		} else { // 换到左边
			r[lt], r[i] = r[i], r[lt]
			lt++
			i++
		}
	}
	if lt-1 != lo {
		r[lt-1], r[lo] = r[lo], r[lt-1]
	}
	lt--
	return lt, gt
}

func _quick_sort(r []rune, lo int, hi int) {
	if lo >= hi { // 一个元素
		return
	}
	lt, gt := _partition(r, lo, hi)
	_quick_sort(r, lo, lt-1)
	_quick_sort(r, gt+1, hi)
}

func sortString(s string) (news string) {
	r := []rune(s) // 将 rune 排序。快速排序吧
	_quick_sort(r, 0, len(r)-1)
	news = string(r)
	return
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "ate", "wocao", "caowo"}))
}

