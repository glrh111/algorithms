package main

// 二叉树实现的 ST

/*
    二叉树的结点
 */
type BSTNode struct {
	key *Comparable
	value interface{}
	leftNode, rightNode *BSTNode
	size int
}

func NewBSTNode(key *Comparable, value interface{}, size int) *BSTNode {
	return &BSTNode{
		key: key,
		value: value,
		size: size,
	}
}

func (this *BSTNode) Size() (size int) {
	if (this == nil) {
		size = 0
	} else {
		size = this.size
	}
	return
}

/*
    二叉树本身
 */
type BST struct {
	root *BSTNode
}

func NewBST() *BST {
	return &BST{
		root: nil,
	}
}

func (this *BST) Size() (size int) {
	if (this.root == nil) {
		size = 0
	} else {
		size = this.root.size
	}
	return
}

func (this *BST) get(node *BSTNode, key *Comparable) (value interface{}) {
	if (nil == node) {
		value = nil
	} else {
		comRe := node.key.CompareTo(*key)
		if (comRe == 0) {         // 命中
			value = node.value
		} else if (comRe < 0 ) {  // 去右边的树上查找
			value = this.get(node.rightNode, key)
		} else {                  // 去左边的树上查找
			value = this.get(node.leftNode, key)
		}
	}
	return
}

// GET
func (this *BST) Get(key *Comparable) (value interface{}) {
	// 从根结点开始查找
	value = this.get(this.root, key)
	return
}

// 如果新建了BSTnode的话，返回出去
// newBst 直接建立的结点  ifCreatNode 路径上是否建立新结点
// 书上的实现是，每次新结点重新计算一下 node.size
func (this *BST) put(node *BSTNode, key *Comparable, value interface{}) (newBst *BSTNode, ifCreatNode bool) {
	//fmt.Println("\n insert key: ", key, value)
	ifCreatNode = false
	if (nil == node) {
		newBst = NewBSTNode(key, value, 1)
		ifCreatNode = true
	} else {
		//fmt.Println("before insert key: ", key, " node.size is: ", node.size)
		compRe := node.key.CompareTo(*key)
		if (0 == compRe) {   // key 存在， 更新 value
			node.value = value
		} else if (compRe < 0) { // 新key 比该结点大，去右侧
			subNewBST, ifSubCreatNode := this.put(node.rightNode, key, value)
			if (subNewBST != nil) {
				node.rightNode = subNewBST
			}
			if ifSubCreatNode {
				node.size++
				ifCreatNode = ifSubCreatNode
			}
		} else {
			subNewBST, ifSubCreatNode := this.put(node.leftNode, key, value)
			if (subNewBST != nil) {
				node.leftNode = subNewBST
			}
			if ifSubCreatNode {
				node.size++
				ifCreatNode = ifSubCreatNode
			}
		}
		//fmt.Println("after insert key: ", key, " node.size is: ", node.size)
	}

	return
}

// PUT
func (this *BST) Put(key *Comparable, value interface{}) {
	re, _ := this.put(this.root, key, value)
	if re != nil {
		this.root = re
	}
}

// Delete 删除
func (this *BST) Delete(key *Comparable) {

}

func (this *BST) contains(node *BSTNode, key *Comparable) (ifContains bool) {
	if (nil == node) {
		ifContains = false
	} else {
		compRe := node.key.CompareTo(*key)
		if (0 == compRe) {   // 找到了
			ifContains = true
		} else if (compRe < 0) { // 在右边查找
			ifContains = this.contains(node.rightNode, key)
		} else {
			ifContains = this.contains(node.leftNode, key)
		}
	}
	return
}

// Contains
func (this *BST) Contains(key *Comparable) bool {
	// 从根上开始查找
	return this.contains(this.root, key)
}

// IsEmpty
func (this *BST) IsEmpty() bool {
	return 0 == this.Size()
}

func (this *BST) keys(node *BSTNode) (ks []*Comparable) {
	if (nil == node) {
		ks = []*Comparable{}
	} else {
		ks = append(
			this.keys(node.leftNode),
			append(
				[]*Comparable{node.key},
				this.keys(node.rightNode)...,
			)...,
		)
	}
	return
}

// Keys
func (this *BST) Keys() (keys []*Comparable) {
	return this.keys(this.root)
}


func readAndCountByBST(filename string, lengthThreshold int) (totalWordCount int, differendWordCount int) {
	return ReadAndCount(
		NewBST(),
		filename,
		lengthThreshold,
	)
}
