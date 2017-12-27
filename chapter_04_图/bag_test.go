// 测试内容
// 每个文件一个测试内容

package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	//"fmt"
	//"fmt"
	"fmt"
)


func TestBag(t *testing.T) {

	// 简单测试
	Convey("功能测试", t, func() {
		fmt.Println("Bag 功能测试。")

		// 首先往里边push数据，看取出来的对不对。
		var bag BagInterface = NewBag(10)

		vList := make([]int, 100000)
		for i:=0; i<100000; i++ {
			if i < 50000 {
				vList[i] = i
			} else {
				vList[i] = i - 50000
			}
		}

		// IsEmpty
		So(bag.IsEmpty(), ShouldEqual, true)

		// Add
		for _, v := range vList {
			bag.Add(NewKey(v))
		}

		// Size
		So(bag.Size(), ShouldEqual, 50000)

		// IsEmpty
		So(bag.IsEmpty(), ShouldEqual, false)

		// iterator
		it := bag.Iterator()
		totalCount := 0
		for {
			_, ok := it()
			if !ok {
				break
			}
			totalCount++
		}
		So(totalCount, ShouldEqual, 50000)

		// chan
		totalCount = 0
		c := bag.IteratorChan()
		for _ = range c {
			totalCount++
			//fmt.Println("v: ", v)
		}
		So(totalCount, ShouldEqual, 50000)

	})

}