package main

import (
	"time"
	//"fmt"
)

/*
   对于每个排序的元素，启动一个 goroutine sleep一段时间后，向channel 中传送本身的值
 */
func SleepSort(data IntSlice) (newData IntSlice) {
	newData = IntSlice([]int{})
	c := make(chan int)
	for _, i := range data {
		go func(j int) {
			time.Sleep(time.Duration(j) * time.Second)
			c <- j
		}(i)
	}
	var (
		l = data.Len()
		count int
	)
	for value := range c {
		newData = append(newData, value)
		count++
		if count == l {
			close(c)
		}
	}
	return
}

//func main() {
//	p := generateIntSlice(1000, 10)
//	fmt.Println("Sleep排序前", p)
//	newp := SleepSort(p)
//	fmt.Println(newp, newp.IsSorted())
//}
