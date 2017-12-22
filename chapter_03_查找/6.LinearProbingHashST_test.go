// 测试内容
// 每个文件一个测试内容

package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"fmt"
	"time"
)


// SYMBOLTABLEEXAMPLE
// 012345678901234567
// SY  O T B   XAMPLE
// 排序过的 ABELMOPSTXY
// 第一次删除过的 ABELMOPSTX
func TestLinearProbingHashST(t *testing.T) {

	// 简单测试
	Convey("测试普通接口", t, func() {
		// 首先往里边push数据，看取出来的对不对。
		var st SymbolTableInterface = NewLinearProbingHashST(97)
		stMap := map[string]int{
			"S": 0,
			"Y": 1,
			"O": 4,
			"T": 6,
			"B": 8,
			"X": 12,
			"A": 13,
			"M": 14,
			"P": 15,
			"L": 16,
			"E": 17,
		}
		// 插入值
		for index, value := range "SYMBOLTABLEEXAMPLE" {
			//fmt.Println("in TestSelfInAppEventReportHandler Put ", index, " ", value)
			st.Put(NewComparable(string(value)), index)
		}
		// 检测
		for key, value := range stMap {
			So(st.Get(NewComparable(key)), ShouldEqual, value)
		}

		// 检测删除
		st.Delete(NewComparable("Y"))
		So(st.Get(NewComparable("Y")), ShouldEqual, nil)
		// 检测 size
		So(st.Size(), ShouldEqual, 10)

		st.Delete(NewComparable("Y"))
		So(st.Get(NewComparable("Y")), ShouldEqual, nil)
		// 检测 size
		So(st.Size(), ShouldEqual, 10)

		// 检测keys

		fmt.Println("keys", st.Keys())

		for _, key := range st.Keys() {
			fmt.Println("key: ", key.Value())
		}

	})

	// 读取著名测试用例
	Convey("测试hashcode", t, func() {

		fmt.Println("In readAndCountByLinearProbingHashST---")
		for _, i := range []int{1, 8, 10} {
			start := time.Now()
			// fixme
			// 读取 tale        116(BST), 146(SCHashST) ms   85ms(这个)(最快，碉堡了)
			// 读取 leipzig1M   (Size通过其他方法实现的版本)32997,9622,4306  BST 31651,9931,4420 ms leipzig1M  16682,5044,2793(真是快, 碉堡了)
			t, d := readAndCountByLinearProbingHashST("exam/leipzig1M.txt", i)
			end := time.Now()
			spendMilliSecond := end.Sub(start).Nanoseconds() / 1000000 // 毫秒
			fmt.Printf("Total words [%d], Different words [%d], Spend MilliSeconds [%d]!\n", t, d, spendMilliSecond)
		}

	})

}