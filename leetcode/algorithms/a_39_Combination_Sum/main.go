package main

import "fmt"

/*
   给定一个数组元素，和一个和的目标。兑换零钱的问题。

   [2, 3, 6, 7]  7

   算法，先排序。然后从大到小递归。

   结束的标志：
      + candidates 为空

 */
func combinationSum(candidates []int, target int) [][]int {
	return combinationsum(candidates, target, []int{})
}

/*
   带有前缀的
 */
func combinationsum(candidates []int, target int, prefix []int) [][]int {

	var (
		clist = [][]int{}
		leftTarget int
	)

	if len(candidates) == 0 {
		return clist
	}

	for i:=0; true; i++ {
		leftTarget = target - i * candidates[0]
		if leftTarget < 0 { // 剩下 负数
			goto outofloop
		} else { // 刚好凑齐
			newcom := make([]int, len(prefix))
			copy(newcom, prefix)
			// 追加 i 个 该 candidates[0] 刚好结束
			for j:=0; j<i; j++ {
				newcom = append(newcom, candidates[0])
			}

			if leftTarget == 0 {
				clist = append(clist, newcom)
				goto outofloop
			} else {
				clist = append(
					clist,
					combinationsum(candidates[1:], leftTarget, newcom)...,
				)
			}

		}

	}

outofloop:
	return clist
}

func main() {
	fmt.Println(combinationSum([]int{1,2,3}, 7))
}
