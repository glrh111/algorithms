package compress

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

/*
    测试字母表
 */

func TestLSDSort(t *testing.T) {

	// 简单测试
	Convey("Alphabet 功能测试", t, func() {

		// ACTG 字母表

		// ToChar(index) rune
		for key, value := range []rune("ACTG") {
			So(DNAALPHABET.ToChar(key), ShouldEqual, value)
		}


	})

}
