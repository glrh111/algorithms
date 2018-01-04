package compress

import (
	//"bytes"
	"io"
	"os"
	"fmt"
)

/*
   读取标准二进制输入流

	可以参考这个项目：
	https://github.com/biogo/biogo

   内部的 bytes 包
 */

type BinaryStdIn struct {
	stdin io.Reader
}

// 返回一个标准
func NewBinaryStdIn() *BinaryStdIn {
	return &BinaryStdIn{
		stdin: os.Stdin,
	}
}

func (si *BinaryStdIn) ReadAndWrite() {
	b := make([]byte, 10) // 缓冲区有 10
	for {
		n, err := si.stdin.Read(b)

		// 处理数据
		fmt.Printf("读取了 %d 个新字节\n", n)
		fmt.Println(b)

		if err != nil {
			break
		}
	}
}





