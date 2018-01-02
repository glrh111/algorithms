package stringst

import (
	"testing"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
)

/*
    低位优先排序
 */

// w 为字符串宽度
func TestLSDSort(t *testing.T) {

	examStringList := []string{
		"she", "sells", "sea", "shells", "by", "the", "sea", "shore", // size = 7
	}

	// 简单测试
	Convey("功能测试", t, func() {

		// 构造函数
		stringst := NewTrieST()

		// PUT
		for index, key := range examStringList {
			stringst.Put(key, index)
		}

		// Size
		So(stringst.Size(), ShouldEqual, 7)

		// GET
		So(stringst.Get("sea").(int), ShouldEqual, 6)

		for index, key := range examStringList {
			fmt.Println("index: ", index, "key: ", stringst.Get(key))
		}

		// CONTAINS
		So(stringst.Contains("sea"), ShouldEqual, true)
		So(stringst.Contains("se0"), ShouldEqual, false)
		So(stringst.Contains("shells"), ShouldEqual, true)

		// DELETE

		// KEYS
		for s := range stringst.Keys() {
			fmt.Println(s)
		}

		// KeysWithPreffix
		for s := range stringst.KeysWithPreffix("wocao") {
			fmt.Println(s)
		}

		for s := range stringst.KeysWithPreffix("se") {
			fmt.Println("wocao", s)
		}

		// LongestPrefixOf
		So(stringst.LongestPrefixOf("sh"), ShouldEqual, "shells")
		So(stringst.LongestPrefixOf("se"), ShouldEqual, "sells")

		for key := range stringst.KeysThatMatch("s..") {
			fmt.Println("match ...: ",  key)
		}

	})

}
