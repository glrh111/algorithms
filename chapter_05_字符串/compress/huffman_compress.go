package compress

import "fmt"

/*
   实现霍夫曼压缩
   编译表： byte -> []bool
 */

// 查找树结点
type HuffmanTreeNode struct {
	left, right *HuffmanTreeNode
	freq int // 出现的频率 , 用 int 表示
	ch byte  // 代表哪个字符
}

func NewHuffmanTreeNode(ch byte, freq int, left *HuffmanTreeNode, right *HuffmanTreeNode) *HuffmanTreeNode {
	return &HuffmanTreeNode{
		left: left,
		right: right,
		freq: freq,
		ch: ch,
	}
}

// 是否是叶子结点。left, right == nil
func (node *HuffmanTreeNode) IsLeaf() bool {
	return node.left == nil && node.right == nil
}

// 霍夫曼单词查找树
type HuffmanTree struct {
	root *HuffmanTreeNode
	compilerTable [][]bool // byte -> []bool 编译表
}

func NewHuffmanTree(root *HuffmanTreeNode, compilerTable [][]bool) HuffmanTree {
	return HuffmanTree{
		root: root,
		compilerTable: compilerTable,
	}
}

func NewHuffmanTreeFromFile(filename string) (tree HuffmanTree) {
	tree = NewHuffmanTree(nil, nil)

	// 读取文件，构造树 + compilerTable
	bsi := NewBinaryStdIn(filename)

	// 频率统计列表
	freqList := make([]int, 256) // ch -> count

	// 第一遍读取
	for {
		ch := bsi.ReadChar()
		if bsi.err != nil {
			break
		}
		freqList[ch] += 1
	}

	// 根据频率统计，构造树. 每一次都得排序. 需要一个 优先队列 数据类型. 能够自动排序的
	pq := NewPriorityQueue()
	for ch, freq := range freqList {
		if freq > 0 {
			pq.Enqueue(freq, NewHuffmanTreeNode(
				byte(ch),
				freq,
				nil, nil,
			)) // value 是 HuffNode 类型，转化时注意
		}
	}

	// 合并结点 每次 Dequeue 频率最低的一个结点
	for pq.Size() >= 2 {
		leftFreq, leftValue, _ := pq.DequeueMin()
		leftNode, _ := leftValue.(*HuffmanTreeNode)
		rightFreq, rightValue, _ := pq.DequeueMin()
		rightNode, _ := rightValue.(*HuffmanTreeNode)

		// 构造新结点
		newNode := NewHuffmanTreeNode(0, rightFreq+leftFreq, leftNode, rightNode)
		pq.Enqueue(newNode.freq, newNode)
	}

	_, rootValue, _ := pq.DequeueMin()
	rootNode, _ := rootValue.(*HuffmanTreeNode)

	//
	tree.root = rootNode

	// 构造编译树
	tree.BuildCompilerTable()

	return tree
}

func (tree *HuffmanTree) ChToBitarr(ch byte) (bits []bool) {
	return tree.compilerTable[ch]
}

/*
    霍夫曼压缩函数 Compress
 */
func HuffmanCompress(fromFilename string, toFilename string, huffmanTreeFilename string) {
	flin := NewBinaryStdIn(fromFilename)
	flout := NewBinaryStdOut(toFilename)

	huffmantree := NewHuffmanTreeFromFile(fromFilename)

	var charAmount uint64

	// 站坑 64 bits
	for i:=0; i<64; i++ {
		flout.WriteBool(false)
	}

	// 实际压缩
	for {
		fromChar := flin.ReadChar()
		if flin.err != nil {
			break
		}
		charAmount++
		compressBitList := huffmantree.ChToBitarr(fromChar)
		for _, bit := range compressBitList {
			flout.WriteBool(bit)
		}
	}

	// FLush 一遍
	flout.Flush(true)

	// 写入 charAmount
	charAmountBit := unsignedToBit(charAmount, 64)
	for i:=0; i<64; i+=8 {
		bs := []byte{bitToByte(charAmountBit[i:i+8])}
		off := int64(i/8)
		flout.WriteAt(bs, off)
	}

	flout.Close()
	WriteHuffmanTree(huffmantree, huffmanTreeFilename)
	// 提示压缩完毕
	fmt.Println("压缩完毕，输出至文件：", toFilename)
}

/*
   霍夫曼解压函数 Expand
 */
func HuffmanExpand(fromFilename string, huffmanTreeFilename string, toFilename string) {
	huffmantree := ReadHuffmanTree(huffmanTreeFilename)
	huffmantree.BuildCompilerTable()

	flin := NewBinaryStdIn(fromFilename)
	flout := NewBinaryStdOut(toFilename)

	var charAmount uint64
	for i:=0; i<8; i++ { // 前八个字节，供字符数量使用。
		c := uint64(flin.ReadChar())
		fmt.Println(i, c, string(c))
		charAmount += (c << uint64((7-i)*8))
	}

	// 按 bit 读数据
	currentNode := huffmantree.root
	var charCount uint64
	for {

		if currentNode.IsLeaf() {
			flout.WriteChar(currentNode.ch)
			currentNode = huffmantree.root
			charCount++
		}

		bit := flin.ReadBit()
		if flin.err != nil || charCount >= charAmount {
			break
		}

		if bit {
			currentNode = currentNode.right
		} else {
			currentNode = currentNode.left
		}

	}
	flout.Close()
	fmt.Println("解压完成，写入文件：", toFilename)
}

/*
    将霍夫曼树写入文件
    前序遍历
    如果是 0,那么不是叶子结点; 如果是 1, 那么是叶子结点，而且后 8 bit为该 char 信息
 */
func walkHuffmanTree(node *HuffmanTreeNode, flout *BinaryStdOut, leafAmount *uint64) {
	if node.IsLeaf() {
		flout.WriteBool(true)
		flout.WriteChar(node.ch)
		*leafAmount++
	} else {
		flout.WriteBool(false)
		walkHuffmanTree(node.left, flout, leafAmount)
		walkHuffmanTree(node.right, flout, leafAmount)
	}
}

func WriteHuffmanTree(tree HuffmanTree, toFilename string) {
	flout := NewBinaryStdOut(toFilename)
	for i:=0; i<64; i++ {
		flout.WriteBool(false)
	}
	var leafAmount uint64
	walkHuffmanTree(tree.root, &flout, &leafAmount)

	// 写入 leaf amount
	flout.Flush(true)

	charAmountBit := unsignedToBit(leafAmount, 64)
	for i:=0; i<64; i+=8 {
		bs := []byte{bitToByte(charAmountBit[i:i+8])}
		off := int64(i/8)
		flout.WriteAt(bs, off)
	}
	flout.Close()
}

/*
   从文件读取一棵霍夫曼树
 */
func readhuffmantree(flin *BinaryStdIn, currentLeafAmount *uint64, totalLeafAmount uint64) (node *HuffmanTreeNode) {
	bit := flin.ReadBit()
	if flin.err != nil {
		return // nil
	}

	if bit { // 1 再次读取后边8位

		ch := flin.ReadChar()
		*currentLeafAmount++
		node = NewHuffmanTreeNode(ch, 0, nil, nil)

	} else {
		node = NewHuffmanTreeNode(0, 0, readhuffmantree(flin, currentLeafAmount, totalLeafAmount),
			readhuffmantree(flin, currentLeafAmount, totalLeafAmount))
	}

	return
}
func ReadHuffmanTree(filename string) (tree HuffmanTree) {
	flin := NewBinaryStdIn(filename)
	// 读取 leafAmount
	var leafAmount uint64
	for i:=0; i<8; i++ { // 前八个字节，供字符数量使用。
		c := uint64(flin.ReadChar())
		leafAmount += (c << uint64((7-i)*8))
	}

	tree = NewHuffmanTree(nil, nil)

	var currentLeafAmount uint64
	tree.root = readhuffmantree(&flin, &currentLeafAmount, leafAmount)
	flin.Close()
	return
}

/*
   从霍夫曼树构造编译表
 */
func (tree *HuffmanTree) buildCT(ct [][]bool, node *HuffmanTreeNode, prelist []bool)  {
	if node != nil {
		if node.IsLeaf() { // 里边含有值
			ct[node.ch] = prelist
		} else {
			if node.left != nil {
				leftCopy := make([]bool, len(prelist))
				copy(leftCopy, prelist)
				tree.buildCT(ct, node.left, append(leftCopy, false)) // 左边是 0
			}
			if node.right != nil {
				rightCopy := make([]bool, len(prelist))
				copy(rightCopy, prelist)
				tree.buildCT(ct, node.right, append(rightCopy, true)) // 右边是 1
			}
		}
	}
}

func (tree *HuffmanTree) BuildCompilerTable() {
	ct := make([][]bool, 256)
	// 构造ct : byte -> []bool
	tree.buildCT(ct, tree.root, []bool{})
	tree.compilerTable = ct
}