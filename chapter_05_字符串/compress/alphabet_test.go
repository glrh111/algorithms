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
		// ToIndex(rune) index
		// Contains(rune) bool
		for key, value := range []rune("ACTGW") {
			if key < 4 {
				So(DNAALPHABET.ToChar(key), ShouldEqual, value)
				So(DNAALPHABET.ToIndex(value), ShouldEqual, key)
				So(DNAALPHABET.Contains(value), ShouldEqual, true)
			} else {
				So(DNAALPHABET.Contains(value), ShouldEqual, false)
			}
		}

		// R
		So(DNAALPHABET.R(), ShouldEqual, 4)

		// LgR
		So(DNAALPHABET.LgR(), ShouldEqual, 2)

		// ToIndices
		ids := DNAALPHABET.ToIndices([]rune("ACTGGATCTG")) //
		expectedIds := []int{0,1,2,3,3,0,2,1,2,3}
		for index, value := range ids {
			So(value, ShouldEqual, expectedIds[index])
		}

		// ToChars
		rl := DNAALPHABET.ToChars(expectedIds) // []rune
		expectedRl := []rune("ACTGGATCTG")
		for index, r := range rl {
			So(r, ShouldEqual, expectedRl[index])
		}

		// New Alphabet
		fca := NewAlphabet("ABCDE")
		So(fca.LgR(), ShouldEqual, 3)

	})

}
