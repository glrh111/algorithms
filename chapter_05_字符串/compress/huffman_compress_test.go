package compress

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

/*
    测试 gene 压缩
 */

func TestHuffmanCompress(t *testing.T) {

	// 简单测试
	Convey("霍夫曼压缩 功能测试", t, func() {
		HuffmanCompress("exam/rawhuff.txt", "exam/huffc.huff", "exam/huff.tree")
		HuffmanExpand("exam/huffc.huff", "exam/huff.tree","exam/huffe.txt")

		//flin := NewBinaryStdIn("exam/huff.tree")
		//for bit := range flin.ReadAndWrite() {
		//	if bit {
		//		fmt.Print("1")
		//	} else {
		//		fmt.Print("0")
		//	}
		//}
		//
		//flin = NewBinaryStdIn("exam/huffc.huff")
		//for bit := range flin.ReadAndWrite() {
		//	if bit {
		//		fmt.Print("1")
		//	} else {
		//		fmt.Print("0")
		//	}
		//}

	})

}
