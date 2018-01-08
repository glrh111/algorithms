package compress

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"fmt"
)

/*
    测试字母表
 */

func TestBinaryStd(t *testing.T) {

	// 简单测试
	Convey("BinaryStdIn 功能测试", t, func() {
		//std := NewBinaryStdIn("alphabet.go")
		//std.ReadAndWrite()
		//std.Close()
		fmt.Println(byteToBit(3))
		fmt.Println(bitToByte([]bool{true, true, false, false, false, false, false, true}))
		fmt.Println(bitarrToBytearr([]bool{true, true, false, false, false, false, false, true}))
		fmt.Println(unsignedToBit(3, 64))

		// chan bool 一次读取一个 char
		//flin := NewBinaryStdIn("exam/rawgene.txt")
		//for bit := range flin.ReadAndWrite() {
		//
		//	ch := 0
		//	if bit {
		//		ch = 1
		//	}
		//
		//	fmt.Print(ch)
		//}
		//
		//// 一次读取一个 bit
		//flin = NewBinaryStdIn("exam/rawgene.txt")
		//for {
		//	c := flin.ReadBit()
		//	if flin.err != nil {
		//		break
		//	}
		//	ch := 0
		//	if c {
		//
		//		ch = 1
		//	}
		//	fmt.Print(ch)
		//}
	})

}
