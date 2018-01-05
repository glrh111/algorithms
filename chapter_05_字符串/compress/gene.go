package compress

import "fmt"

func GeneCompress(fromFilename string, toFilename string) {
	// 从前者读取数据，写入到后者里边. 后者文件的 后 64 bit(8 byte), 代表文件含有多少字母。用作解析
	flin := NewBinaryStdIn(fromFilename)
	flout := NewBinaryStdOut(toFilename)

	// n
	var charAmount uint64

	// 站坑
	for i:=0; i<64; i++ {
		flout.WriteBool(false)
	}

	for {
		c := flin.ReadChar()
		if flin.err != nil {
			break
		}
		cIndex := DNAALPHABET.ToIndex(rune(c)) // 将 cIndex 写入里边 [0, 3] 需要两位即可表示

		// [ 高位, 低位 ]
		flout.WriteBool(cIndex & 10 != 0)
		flout.WriteBool(cIndex & 1 != 0)
		charAmount += 1
	}

	flout.Flush(true) // 将全部字节 写入一遍，然后写入 总 char 数量

	// 将 字符数量，写入到文件前面
	charAmountBit := unsignedToBit(charAmount, 64)
	for i:=0; i<64; i+=8 {
		bs := []byte{bitToByte(charAmountBit[i:i+8])}
		off := int64(i/8)

		n, err := flout.WriteAt(bs, off)
		fmt.Println(bs, off, n, err)
	}

	fmt.Println("写入的char数量为：", charAmount)

	flout.Close()

	fmt.Println("压缩完成!")

}

func GeneExpand(fromFilename string, toFilename string) {
	// 认真处理最后一个字节
	flin := NewBinaryStdIn(fromFilename)
	//flout := NewBinaryStdOut(toFilename)

	// 解析出来字符数量
	var charAmount uint64
	for i:=0; i<8; i++ { // 前八个字节，供字符数量使用。
		c := uint64(flin.ReadChar())
		charAmount += (c << uint64((7-i)*8))
	}

	fmt.Println("解析出来的char数量为：", charAmount)

}
