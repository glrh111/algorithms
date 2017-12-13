// 测试内容
// 每个文件一个测试内容

package main

import (
	. "github.com/smartystreets/goconvey/convey"
	//"time"
	"testing"
	//"fmt"
	"fmt"
)

// 获取测试结果。网上公开的一些测试用例
func getTestResult() {

}

// SYMBOLTABLEEXAMPLE
// 012345678901234567
// SY  O T B   XAMPLE
func TestSelfInAppEventReportHandler(t *testing.T) {

	Convey("测试post接口", t, func() {
		// 首先往里边push数据，看取出来的对不对。
		var st SymbolTableInterface = NewLinkedList()
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
		// 检测size
		So(st.Size(), ShouldEqual, 10)
		// 检测keys

		keyValueList := []string{}
		for _, key := range st.Keys() {
			fmt.Println("key, ", key, *key, key.value, key.Value())
			keyValueList = append(keyValueList, "s")
		}

		for index, _ := range stMap {
			if ("Y" == index) { continue }
		}
	})


}