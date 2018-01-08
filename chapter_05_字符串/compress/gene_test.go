package compress

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

/*
    测试 gene 压缩
 */

func TestGene(t *testing.T) {

	// 简单测试
	Convey("Gene 压缩解压 功能测试", t, func() {
		var (
			rawFilename = "exam/rawgene.txt"
			c1filename = "exam/c1.gene" // 第一次压缩的文件名字
			e1filename = "exam/e1.txt"  // 第一次解压的文件名
		)

		GeneCompress(rawFilename, c1filename)

		GeneExpand(c1filename, e1filename)

	})

}
