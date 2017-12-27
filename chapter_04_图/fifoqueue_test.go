// 测试内容
// 每个文件一个测试内容

package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"fmt"
)


func TestFIFOQueue(t *testing.T) {

	// 简单测试
	Convey("功能测试", t, func() {
		fmt.Println("FIFOQueue 功能测试。")

		// 首先往里边push数据，看取出来的对不对。
		var queue FIFOQueueInterface = NewFIFOQueue()

		// Size
		So(queue.Size(), ShouldEqual, 0)

		for i:=0; i<500; i++ {
			queue.Enqueue(i)
		}

		So(queue.Size(), ShouldEqual, 500)

		// 测试 dequeue
		for i:=0; i<500; i++ {
			value, ok := queue.Dequeue()
			So(value, ShouldEqual, i)
			So(ok, ShouldEqual, true)
		}

		// Size
		So(queue.Size(), ShouldEqual, 0)

		_, ok := queue.Dequeue()
		So(ok, ShouldNotEqual, true)

	})

}