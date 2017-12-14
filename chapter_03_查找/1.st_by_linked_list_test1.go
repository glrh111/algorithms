// 测试内容
// 每个文件一个测试内容

package main

import (
	. "github.com/smartystreets/goconvey/convey"
	//"time"
	"testing"
	//"fmt"
	"fmt"
	"time"
)

// 获取测试结果。网上公开的一些测试用例
// tinyTale 双城记前5行
// tale     双城记全书
// leipzig corpora collection leipzig1M.txt 网上随机选取的1M句子。
func getTestResult() interface{} {
	type fileInfo struct {
		WordCount int
		DifferentWordCount int
		WordCount8 int          // 长度大于等于10的单词
		DifferentWordCount8 int // 长度大于等于10的单词数
		WordCount10 int
		DifferentWordCount10 int
	}
	return map[string]fileInfo{
		"tinyTale": fileInfo{
			60,
			20,
			3,
			3,
			2,
			2,
		},
		"tale": fileInfo{
			135635,
			10679,
			14350,
			5737,
			4582,
			2260,
		},
		"leipzig1M": fileInfo{
			21191455,
			534580,
			4239597,
			299593,
			1610829,
			165555,
		},
	}
}

// SYMBOLTABLEEXAMPLE
// 012345678901234567
// SY  O T B   XAMPLE
func TestLinkedListST(t *testing.T) {

	// 简单测试
	Convey("测试普通接口", t, func() {
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
			keyValueList = append(keyValueList, key.Value().(string))
		}

		for index, _ := range stMap {
			if ("Y" == index) { continue }
		}
	})

	// 读取著名测试用例
	Convey("测试著名测试用例", t, func() {
		// 读取文件内容
		for _, i := range []int{1, 8, 10} {
			start := time.Now()
			t, d := readAndCount("exam/tale.txt", i)
			end := time.Now()
			spendMilliSecond := end.Sub(start).Nanoseconds() / 1000000 // 秒
			fmt.Printf("Total words [%d], Different words [%d], Spend MilliSeconds [%d]!\n", t, d, spendMilliSecond)
		}
	})

}