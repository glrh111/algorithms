package compress

import (
	"io"
	"os"
	"fmt"
	"bufio"
)

/*
   读取标准二进制输入流

	可以参考这个项目：
	https://github.com/biogo/biogo

   内部的 bytes 包
 */

type BinaryStdIn struct {
	BitReader
	fl *os.File
}

// 返回一个标准
func NewBinaryStdIn(filename string) BinaryStdIn {
	fl, err := os.Open(filename)
	if err != nil {
		panic("Open file error: " + err.Error())
	}
	return BinaryStdIn{NewBitReader(fl),fl}
}

func (si *BinaryStdIn) ReadBool() bool {
	return si.ReadBit()
}

// FIXME 本函数与ReadBool 混合使用，会产生问题。
func (si *BinaryStdIn) ReadChar() byte {
	//if way {
	//	b, err := si.r.ReadByte()
	//	si.err = err
	//	return b
	//} else {
		bits := make([]bool, 8)
		for i := 0; i < 8; i++ {
			bits[i] = si.ReadBool()
		}
		return bitToByte(bits)
	//}
}

// 读完了之后，就是empty
func (si *BinaryStdIn) IsEmpty() bool {
	return si.err != nil
}

func (si *BinaryStdIn) Close() {
	// 关闭文件
	si.fl.Close()
}


func (si *BinaryStdIn) ReadAndWrite() (c chan bool) {
	c = make(chan bool)
	chAmount := 0 // 文件 bit 大小
	go func() {

		for {
			ch := si.ReadChar() // byte
			if si.err != nil {  // 读取到文件结尾了
				fmt.Println("文件bit: ", chAmount)
				close(c)
				break
			}
			for _, bit := range byteToBit(ch) { // 高位 - 低位
				c <- bit
				chAmount += 1
			}
		}
	}()

	return
}


/*
    BinaryStdOut
    这个不需要 1 bit来写，操作系统也不直接支持。先蹿在缓冲区里边，最后一块儿写入
 */
type BinaryStdOut struct {
	fl *os.File
	buffsize int // 缓冲区里面有多少 byte ？
	bitbuff []bool // 缓冲区，bit。写入之前检查下是否为整数字节
}

func NewBinaryStdOut(filename string) BinaryStdOut {
	os.Remove(filename)
	fl, err := os.Create(filename)
	if err != nil {
		panic("os error: " + err.Error())
	}
	return BinaryStdOut{fl, 9192, []bool{}} // 1K buff
}

func (so *BinaryStdOut) WriteBool(bit bool) {
	so.bitbuff = append(so.bitbuff, bit)
	so.Flush(false) // 效率有点低，每次都得比较. 不过先这么来吧
}

func (so *BinaryStdOut) WriteChar(b byte) {
	// 先转换为 []bool
	so.bitbuff = append(so.bitbuff, byteToBit(b)...)

	so.Flush(false)
}

// flush 缓冲区内容，到文件里面
// com 代表是否补齐 byte
func (so *BinaryStdOut) Flush(com bool) {
	lenbuff := len(so.bitbuff)
	if lenbuff >= so.buffsize || com {
		bl := len(so.bitbuff) / 8 // 有多少个对齐的字节
		// 查看末尾有多少个字节
		buffend := so.bitbuff[bl*8:]
		if len(buffend)>0 && com { // 补齐最后的字节
			so.bitbuff = append(so.bitbuff, make([]bool, 8-len(buffend))...)
			bl += 1
			buffend = []bool{}
		}
		so.fl.Write(bitarrToBytearr(so.bitbuff)) // 实际写入. 如果没有补齐, 这个函数会忽略最后的几位
		so.bitbuff = buffend
	}
}

func (so *BinaryStdOut) WriteAt(b []byte, off int64) (int, error) {
	return so.fl.WriteAt(b, off)
}

func (so *BinaryStdOut) Close() {
	// Close 之前 Flush 一下。
	so.Flush(true)
	so.fl.Close()
}

/*
   BitReader
 */
type BitReader struct {
	r    io.ByteReader
	n    uint64
	bits uint
	err  error
}

func NewBitReader(r io.Reader) BitReader {
	byter, ok := r.(io.ByteReader)
	if !ok {
		byter = bufio.NewReader(r)
	}
	return BitReader{r: byter}
}

// ReadBits64 reads the given number of bits and returns them in the
// least-significant part of a uint64. In the event of an error, it returns 0
// and the error can be obtained by calling Err().
func (br *BitReader) ReadBits64(bits uint) (n uint64) {
	for bits > br.bits {
		b, err := br.r.ReadByte()
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
		if err != nil {
			br.err = err
			return 0
		}
		br.n <<= 8
		br.n |= uint64(b)
		br.bits += 8
	}

	// br.n looks like this (assuming that br.bits = 14 and bits = 6):
	// Bit: 111111
	//      5432109876543210
	//
	//         (6 bits, the desired output)
	//        |-----|
	//        V     V
	//      0101101101001110
	//        ^            ^
	//        |------------|
	//           br.bits (num valid bits)
	//
	// This the next line right shifts the desired bits into the
	// least-significant places and masks off anything above.
	n = (br.n >> (br.bits - bits)) & ((1 << bits) - 1)
	br.bits -= bits
	return
}

func (br *BitReader) ReadBits(bits uint) (n int) {
	n64 := br.ReadBits64(bits)
	return int(n64)
}

func (br *BitReader) ReadBit() bool {
	n := br.ReadBits(1)
	return n != 0
}

func byteToBit(b byte) (bits []bool) {
	return unsignedToBit(uint64(b), 8)
}

func bitToByte(bits []bool) (b byte) {
	return uint8(bitToUnsigned(bits, 8))
}

// 将 bit 转化为 uint64. n 为 8, 16, 32, 64 分别表示相应位数的无符号数
func bitToUnsigned(bits []bool, n int) (a uint64) {
	for i:=0; i<8; i++ {
		if bits[i] {
			a += 1 << uint64(n-1-i)
		}
	}
	return
}

// 将一个无符号数，转换为 []bit
func unsignedToBit(a uint64, n int) (bits []bool) {
	bits = make([]bool, n) // 高位 --- 低位
	for bit,mask:=0,uint64(1); bit<n; bit,mask=bit+1,mask<<1 {
		bits[n-1-bit] = mask&a != 0
	}
	return
}

// 将超过一个字节的 bit 转换为 []byte 字节流
func bitarrToBytearr(bits []bool) (b []byte) {
	b = make([]byte, len(bits)/8) // 转化为多少字节
	for i:=0; i<len(b); i+=1{
		b[i] = bitToByte(bits[i*8:i*8+8])
	}
	return
}







