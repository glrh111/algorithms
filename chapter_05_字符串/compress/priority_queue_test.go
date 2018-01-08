package compress

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"fmt"
)

/*
    测试 gene 压缩
 */

func TestPriorityQueue(t *testing.T) {

	// 简单测试
	Convey("PriorityQueue 功能测试", t, func() {
		fmt.Println("PriorityQueue 功能测试")
		pq := NewPriorityQueue()
		for value, freq := range []int{1,5,7,8,3,4} {
			pq.Enqueue(freq, value)
		} // 出队列顺序应该是：3,2,1,5,4,0
		for freq, value, ok := pq.DequeueMax(); ok; freq, value, ok = pq.DequeueMax() {
			fmt.Println(freq, value, pq.Size())
		}
		for value, freq := range []int{1,5,7,8,3,4} {
			pq.Enqueue(freq, value)
		} // 出队列顺序应该是：3,2,1,5,4,0
		for freq, value, ok := pq.DequeueMin(); ok; freq, value, ok = pq.DequeueMin() {
			fmt.Println(freq, value, pq.Size())
		}
	})

}
