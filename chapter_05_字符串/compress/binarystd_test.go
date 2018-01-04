package compress

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

/*
    测试字母表
 */

func TestBinaryStd(t *testing.T) {

	// 简单测试
	Convey("BinaryStdIn 功能测试", t, func() {
		std := NewBinaryStdIn()
		std.ReadAndWrite()
	})

}
