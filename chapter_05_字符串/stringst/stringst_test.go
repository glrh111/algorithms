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

		for _, key := range examStringList {
			So(stringst.Contains(key), ShouldEqual, true)
		}

		// KEYS
		for s := range stringst.Keys() {
			fmt.Println(s)
		}

		// KeysWithPreffix
		for s := range stringst.KeysWithPrefix("wocao") {
			fmt.Println(s)
		}

		for s := range stringst.KeysWithPrefix("se") {
			fmt.Println("wocao", s)
		}

		// LongestPrefixOf
		So(stringst.LongestPrefixOf("shellsa"), ShouldEqual, "shells")
		So(stringst.LongestPrefixOf("sean"), ShouldEqual, "sea")
		So(stringst.LongestPrefixOf("wocao"), ShouldEqual, "")

		for key := range stringst.KeysThatMatch("s..") {
			fmt.Println("match ...: ",  key)
		}

		fmt.Println(stringst.Get("she"))

		// DELETE  "she", "sells", "sea", "shells", "by", "the", "sea", "shore"
		ifDelete := stringst.Delete("she")
		So(ifDelete, ShouldEqual, true)
		ifDelete = stringst.Delete("she")
		So(ifDelete, ShouldEqual, false)
		So(stringst.Size(), ShouldEqual, 6)

		So(stringst.Get("shells"), ShouldNotEqual, nil)


	})

}
