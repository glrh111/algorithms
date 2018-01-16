package main

//import "fmt"

/*
   插入排序

   从i=1 开始，将元素插入已经排定的数组中。
   [0, i) 是已经排定的元素
   每次循环的目的，是将 i 插入 [0, i) 中
 */
func InsertionSort(data Interface) {
	for i:=1; i<data.Len(); i++ {
		for j:=i; j>0 && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

/*
   希尔排序  插入排序的改进版本，快的一笔
 */
func ShellSort(data Interface) {
	var (
		l, h = data.Len(), 1
	)

	for {
		if h >= l / 3 {
			break
		}
		h = h * 3 + 1
	}

	for {

		if h < 1 {
			break
		}
		// hi数组变为 h有序
		for i:=h; i<l; i++ {
			for j:=i; j>=h && data.Less(j, j-h); j-=h {
				data.Swap(j, j-h)
			}
		}
		h = h / 3
	}
}


//func main() {
//	p := generateIntSlice(100000, 10)
//	fmt.Println("插入排序前")
//	//InsertionSort(p)
//	//fmt.Println("插入排序后：", p.IsSorted())
//
//	p.Shuffle()
//	ShellSort(p)
//	fmt.Println("希尔排序后：", p.IsSorted())
//}
