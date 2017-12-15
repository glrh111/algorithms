// 测试内容
// 每个文件一个测试内容

package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
	"fmt"
)


// SYMBOLTABLEEXAMPLE
// 012345678901234567
// SY  O T B   XAMPLE
// 排序过的 ABELMOPSTXY
// 第一次删除过的 ABELMOPSTX
func TestBinarySearch(t *testing.T) {

	// 简单测试
	Convey("测试基本内容", t, func() {
		// 首先往里边push数据，看取出来的对不对。
		var st SortedSymbolTableInterface = NewBinarySearchST()
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

		//for _, value := range st.Keys() {
		//	fmt.Println(value.Value())
		//}
		//fmt.Println(len(st.Keys()))

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

		// 测试 Min
		minKey, err := st.Min()
		So(err, ShouldEqual, nil)
		So(minKey.Value(), ShouldEqual, NewComparable("A").Value())

		// 测试空ST MIN
		var stEmpty SortedSymbolTableInterface = NewBinarySearchST()
		minKey, err = stEmpty.Min()
		So(err, ShouldNotEqual, nil)

		// 测试 Max
		maxKey, err := st.Max()
		So(err, ShouldEqual, nil)
		So(maxKey.Value(), ShouldEqual, NewComparable("X").Value())

		// 测试空ST Max
		maxKey, err = stEmpty.Min()
		So(err, ShouldNotEqual, nil)

		// 测试 Contains 在keys里边的，都应该为true
		for _, key := range st.Keys() {
			contains := st.Contains(key)
			So(contains, ShouldEqual, true)
		}

		contains := st.Contains(NewComparable("Y"))
		So(contains, ShouldEqual, false)

		// 测试 Floor 1) 用小于最小的测试 2) 用存在的测试 3) 用大于最大的测试
		floorKey, err := st.Floor(NewComparable("1")) // 应该返回 err
		So(err, ShouldNotEqual, nil)

		floorKey, err = st.Floor(NewComparable("A"))
		So(err, ShouldEqual, nil)
		So(floorKey.Value(), ShouldEqual, "A")

		floorKey, err = st.Floor(NewComparable("Y"))
		So(err, ShouldEqual, nil)
		So(floorKey.Value(), ShouldEqual, "X")

		// 测试 ceiling 1) 小于最小的测试 2) 存在的测试 3) 大于最大的测试
		ceilingKey, err := st.Ceiling(NewComparable("1")) // 应该返回 err
		So(err, ShouldEqual, nil)
		So(ceilingKey.Value(), ShouldEqual, "A")

		ceilingKey, err = st.Ceiling(NewComparable("A"))
		So(err, ShouldEqual, nil)
		So(ceilingKey.Value(), ShouldEqual, "A")

		ceilingKey, err = st.Ceiling(NewComparable("Y"))
		So(err, ShouldNotEqual, nil)

		// 测试 SelectKey 1) 边界的两个值 2) 超出边界的一个值 3) 中间的某个值
		selectedKey, err := st.SelectKey(0)
		So(selectedKey.Value(), ShouldEqual, st.Keys()[0].Value())
		So(err, ShouldEqual, nil)

		selectedKey, err = st.SelectKey(st.Size()-1)
		So(selectedKey.Value(), ShouldEqual, st.Keys()[st.Size()-1].Value())
		So(err, ShouldEqual, nil)

		selectedKey, err = st.SelectKey(st.Size())
		So(err, ShouldNotEqual, nil)

		selectedKey, err = st.SelectKey(2)
		So(selectedKey.Value(), ShouldEqual, "E")
		So(err, ShouldEqual, nil)

		// 测试 DeleteMin 1) 测试存在的 2) 测试空的
		// Now ABELMOPSTX 将会删除 A
		err = st.DeleteMin()
		So(err, ShouldEqual, nil)
		minKey, err = st.Min()
		So(minKey.Value(), ShouldEqual, "B")

		err = stEmpty.DeleteMin()
		So(err, ShouldNotEqual, nil)

		// 测试 DeleteMax 1) 测试存在的 2) 测试空的
		// Now BELMOPSTX 将会删除 X
		err = st.DeleteMax()
		So(err, ShouldEqual, nil)
		maxKey, err = st.Max()
		So(maxKey.Value(), ShouldEqual, "T")

		err = stEmpty.DeleteMax()
		So(err, ShouldNotEqual, nil)

		// SizeBetween 1) 测试两个key都存在于st中的 2) 测试一个key在st中的 3) 两个都不在 4) 测试 hi < lo 的
		// 现在key列表如下 BELMOPST
		size := st.SizeBetween(NewComparable("B"), NewComparable("M")) // ==4
		So(size, ShouldEqual, 4)

		size = st.SizeBetween(NewComparable("A"), NewComparable("M")) // ==4
		So(size, ShouldEqual, 4)

		size = st.SizeBetween(NewComparable("B"), NewComparable("N")) // ==4
		So(size, ShouldEqual, 4)

		size = st.SizeBetween(NewComparable("A"), NewComparable("Y")) // ==8
		So(size, ShouldEqual, 8)

		size = st.SizeBetween(NewComparable("T"), NewComparable("A")) // ==0
		So(size, ShouldEqual, 0)

		// KeysBetween 1) 测试两个key都存在于st中的 2) 测试一个key在st中的 3) 两个都不在 4) 测试 hi < lo 的
		keys := st.KeysBetween(NewComparable("B"), NewComparable("M")) // BELM
		So(len(keys), ShouldEqual, 4)
		for index, key := range keys {
			So(key.Value(), ShouldEqual, string("BELM"[index]))
		}

		keys = st.KeysBetween(NewComparable("B"), NewComparable("N")) // BELM
		for index, key := range keys {
			So(key.Value(), ShouldEqual, string("BELM"[index]))
		}

		keys = st.KeysBetween(NewComparable("T"), NewComparable("A")) // ==0
		So(len(keys), ShouldEqual, 0)

	})



	// 读取著名测试用例
	Convey("测试著名测试用例", t, func() {
		// 读取文件内容
		for _, i := range []int{1, 8, 10} {
			start := time.Now()
			// fixme
			// 读取 tale 1600 ms
			// 读取 leipzig1M 跑了20多分钟，没有跑出来
			t, d := readAndCountByBinarySearchST("exam/tinyTale.txt", i)
			end := time.Now()
			spendMilliSecond := end.Sub(start).Nanoseconds() / 1000000 // 秒
			fmt.Printf("Total words [%d], Different words [%d], Spend MilliSeconds [%d]!\n", t, d, spendMilliSecond)
		}
	})



}