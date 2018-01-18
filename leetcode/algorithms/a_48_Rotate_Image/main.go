package main

import (
	"fmt"
	"math"
)

/*
   给定一个图片矩阵，旋转它
   n * n 的， 要求 in-place 不占用额外空间

   在下面的需要额外空间的基础上，让两个两个交换。存储前一个交换的位置，即可
   但是在实际操作中，发现一次旋转顾及不到所有的元素。所以得多转几次。
 */
func rotate(matrix [][]int)  {

	var (
		lm = len(matrix)
		ceillm2 = int(math.Ceil(float64(lm)/2)) - 1
	)

	// 照顾中心的
	for j:=0; j<=ceillm2; j++ {
		// 横向转
		for i:=j; i<lm-j-1; i++ {
			//fmt.Println("j, i: ", j, i)
			// 起点坐标是 (j, i)
			value_0 := matrix[j][i]
			rotate_su(matrix, j, i, j, i)
			// 将 0, 0 的值，copy到该去的位置
			matrix[i][lm-1-j] = value_0
		}
	}
}

func rotate_su(matrix [][]int, i int, j int, start_i int, start_j int) {
	// 执行交换

	var (
		to_i, to_j = j, len(matrix)-1-i
	)
	//fmt.Printf("(%d, %d) <- (%d, %d)\n", to_i, to_j, i, j)
	if to_i == start_i && to_j == start_j {
		matrix[to_i][to_j] = matrix[i][j]
		return
	}
	rotate_su(matrix, to_i, to_j, start_i, start_j) // 之后应该有操作啊
	matrix[to_i][to_j] = matrix[i][j]
}


/*
    旋转，不需要额外空间的非递归版本。因为递归有额外开销，不如使用循环实现
 */
func rotate_no_re(matrix [][]int)  {

	var (
		lm = len(matrix)
		ceillm2 = int(math.Ceil(float64(lm)/2)) - 1
	)

	// 照顾中心的
	for i:=0; i<=ceillm2; i++ {
		// 横向转
		for j:=i; j<lm-i-1; j++ {
			//fmt.Println("j, i: ", j, i)
			// 起点坐标是 (i, j)
			var (
				value_0 = matrix[i][j] // 起点的值
				start_i, start_j = i, j
				from_i, from_j = start_i, start_j
			)
			for {
				from_i, from_j = lm-1-from_j, from_i
				if from_i == start_i && from_j == start_j {
					matrix[from_j][lm-1-from_i] = value_0
					break
				}
				matrix[from_j][lm-1-from_i] = matrix[from_i][from_j]
			}
		}
	}
}

// 下面的是额外分配空间的版本  坐标系转换  (x, y) -> (y, n-1-x)
// 因为每个点的位置，都要发生变换，所以每次更换位置的时候，被占用位置的点的新位置，可以计算出来。这是一个循环
func rotate_no(matrix [][]int) {
	var (
		lm = len(matrix)
		m = make([][]int, lm)
	)

	for i:=0; i<lm; i++ {
		m[i] = make([]int, lm)
	}

	for i:=0; i<lm; i++ {
		for j:=0; j<lm; j++ {
			m[j][lm-1-i] = matrix[i][j]
		}
	}
	copy(matrix, m)
}

func main() {
	ma := [][]int{
		[]int{1,2,3},
		[]int{4,5,6},
		[]int{7,8,9},
	}
	//ma2 := [][]int{
	//	[]int{0,1},
	//	[]int{2,3},
	//}
	fmt.Println(ma)
	rotate_no_re(ma)
	fmt.Println(ma)


}
