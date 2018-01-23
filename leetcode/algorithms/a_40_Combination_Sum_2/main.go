package main

import "fmt"

/*
   给定一个正整数列表和一个目标和，里边可能有重复元素。求所有元素加和等于这个目标和的组合列表

   结合 39 考虑。直接用39的方法，会导致有重复组合。所以先将列表去重计数。这个树木，就是那个元素所能使用的最大次数
 */
func combinationSum2(candidates []int, target int) [][]int {
	// 求 elemCounter
	var (
		elemCounter = make(map[int]int)
		nodup = []int{}
	)
	for i:=0; i<len(candidates); i++ {
		elemCounter[candidates[i]] += 1
	}
	for key, _ := range elemCounter {
		nodup = append(nodup, key)
	}

	return combinationsum2(nodup, target, []int{}, elemCounter)

}

func combinationsum2(candidates []int, target int, prefix []int, elemCounter map[int]int) (clist [][]int) {
	clist = [][]int{}

	if len(candidates) == 0 {
		return
	}

	for i:=0; i<=elemCounter[candidates[0]]; i++ { // 最多出现这么多次
		leftTarget := target - i * candidates[0]
		if leftTarget < 0 {
			goto outofloop
		} else {
			newcom := make([]int, len(prefix))
			copy(newcom, prefix)
			for j:=0; j<i; j++ { // 追加 i 个值
				newcom = append(newcom, candidates[0])
			}
			if leftTarget == 0 { // 刚好
				clist = append(clist, newcom)
				goto outofloop
			} else {
				clist = append(
					clist,
					combinationsum2(candidates[1:], leftTarget, newcom, elemCounter)...,
				)
			}
		}
	}
outofloop:
	return
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
