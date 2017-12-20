package main

import (
	"errors"
)

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
func (this *BST) put(node *BSTNode, key *Comparable, value interface{}) (newBst *BSTNode) {
	//fmt.Println("\n insert key: ", key, value)
	if (nil == node) {
		newBst = NewBSTNode(key, value, 1)
	} else {
		//fmt.Println("before insert key: ", key, " node.size is: ", node.size)
		compRe := node.key.CompareTo(*key)
		if (0 == compRe) {   // key 存在， 更新 value
			node.value = value
		} else if (compRe < 0) { // 新key 比该结点大，去右侧
			node.rightNode = this.put(node.rightNode, key, value)
		} else {
			node.leftNode = this.put(node.leftNode, key, value)
		}
		node.size = node.leftNode.Size() + node.rightNode.Size() + 1
		newBst = node
	}
	return
}

// PUT
func (this *BST) Put(key *Comparable, value interface{}) {
	this.root = this.put(this.root, key, value)
}

// 是否成功删除，是否需要替换子结点，替换子结点的值 FIXME 参考P261 的简单实现
func (this *BST) delete(node *BSTNode, key *Comparable) (repNode *BSTNode) {

	if node != nil {
		reComp := node.key.CompareTo(*key)

		if reComp == 0 { // 在这里删除. 区分左右侧是否有值

			if node.leftNode == nil {
				node = node.rightNode
			} else if node.rightNode == nil {
				node = node.leftNode
			} else {
				rightMin := this.minNode(node.rightNode)
				rightMin.rightNode = this.deleteMin(node.rightNode) // 删除右侧最小的结点
				rightMin.leftNode = node.leftNode
				node = rightMin
			}
		} else if reComp < 0 { // right
			node.rightNode = this.delete(node.rightNode, key)
		} else {
			node.leftNode = this.delete(node.leftNode, key)
		}
		// 应该返回新结点
		if node != nil {
			node.size = node.leftNode.Size() + node.rightNode.Size() + 1
		}
		repNode = node
	}
	return
}

// Delete 删除
// 参见书本 page 260(总页码) T.Hibbard 方法
func (this *BST) Delete(key *Comparable) {
	this.root = this.delete(this.root, key)
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

// 使用keys 数组，来回传递
func (this *BST) keys(node *BSTNode, keyList *[]*Comparable) {
	if (nil != node) {
		this.keys(node.leftNode, keyList)
		*keyList = append(*keyList, node.key)
		this.keys(node.rightNode, keyList)
	}
}

// Keys
func (this *BST) Keys() (keyList []*Comparable) {
	keyList = []*Comparable{}
	this.keys(this.root, &keyList)
	return
}

func (this *BST) minNode(node *BSTNode) (minN *BSTNode) {
	if nil == node { // root 才会走到这里
		minN = nil
	} else if (node.leftNode == nil) {
		minN = node
	} else {
		minN = this.minNode(node.leftNode)
	}
	return
}

func (this *BST) min(node *BSTNode) (key *Comparable) {
	minNode := this.minNode(node)
	if nil != minNode {
		key = minNode.key  // FIXME : 为了delete需要，选择不复制key
	}
	return
}

// Min 返回最小键，顺着左树查找
func (this *BST) Min() (key *Comparable, err error) {
	key = this.min(this.root)
	if nil == key {
		err = errors.New("BST is empty")
	}
	return
}

func (this *BST) max(node *BSTNode) (key *Comparable) {
	if nil == node { // root 才会走到这里
		key = nil
	} else if (node.rightNode == nil) {
		key = &Comparable{}
		*key = *node.key // 复制一份
	} else {
		key = this.max(node.rightNode)
	}
	return
}

// Max 返回最小键，顺着左树查找
func (this *BST) Max() (key *Comparable, err error) {
	key = this.max(this.root)
	if nil == key {
		err = errors.New("BST is empty")
	}
	return
}

// 如果没找到，返回nil
func (this *BST) floor(node *BSTNode, key *Comparable) (reKey *Comparable) {
	if node == nil {
		reKey = nil
	} else {
		// 跟本结点比较大小
		reComp := node.key.CompareTo(*key)
		if reComp == 0 {         // 返回这个值
			reKey = node.key
		} else if (reComp < 0) { // 试探右树上查找
			reKey = this.floor(node.rightNode, key) // 判断返回值是否为nil
			if reKey == nil {
				reKey = node.key
			}
		} else {                 // 在左树上查找
			reKey = this.floor(node.leftNode, key)
		}
	}
	if nil != reKey {  // 复制一份
		*reKey = *reKey
	}
	return
}

// Floor 小于等于 key 的最大键; 如果没找到，返回 err
func (this *BST) Floor(key *Comparable) (reKey *Comparable, err error) {
	reKey = this.floor(this.root, key)
	if reKey == nil {
		err = errors.New("No return key")
	}
	return
}

// ceiling 没找到返回 nil
func (this *BST) ceiling(node *BSTNode, key *Comparable) (reKey *Comparable) {
	if node == nil {
		reKey = nil
	} else {
		// 跟本结点比较大小
		reComp := node.key.CompareTo(*key)
		if reComp == 0 {   // 返回这个值
			reKey = node.key
		} else if (reComp < 0) { // 试探右树上查找
			reKey = this.ceiling(node.rightNode, key) // 判断返回值 nil
		} else {                 // 在左树上查找
			reKey = this.ceiling(node.leftNode, key)
			if nil == reKey {     // 判断返回值是否为 nil
				reKey = node.key
			}
		}
	}
	if nil != reKey {  // 复制一份
		*reKey = *reKey
	}
	return
}

// Ceiling 大于等于key的最小值; 如果没找到，返回err
func (this *BST) Ceiling(key *Comparable) (reKey *Comparable, err error) {
	reKey = this.ceiling(this.root, key)
	if reKey == nil {
		err = errors.New("No return key")
	}
	return
}

// 应该不会返回nil
func (this *BST) selectKey(node *BSTNode, index int) (key *Comparable) {
	// 不用处理 结点为nil的情况
	// 从左结点的size开始比较
	leftSize := node.leftNode.Size()
	if (leftSize < index) { // 右边树上找
		key = this.selectKey(node.rightNode, index - leftSize - 1)
	} else if (leftSize == index) { // 返回该结点的key
		key = node.key
	} else {                // 在左边树上找
		key = this.selectKey(node.leftNode, index)
	}
	*key = *key // 复制一份
	return
}

// SelectKey
func (this *BST) SelectKey(index int) (key *Comparable, err error) {
	// 检查index范围
	if (index < 0 || index >= this.Size()) {
		err = errors.New("Index out of BST!")
	} else {
		key = this.selectKey(this.root, index)
	}
	return
}

// 需要判断 nil
func (this *BST) rank(node *BSTNode, key *Comparable) (rank int) {
	if nil != node {
		reComp := node.key.CompareTo(*key)
		if (reComp == 0) { // 返回在这棵树上的位置 1 代表该结点
			rank = node.leftNode.Size()
		} else if (reComp < 0) { // 去右边树
			rank = node.leftNode.Size() + 1 + this.rank(node.rightNode, key)
		} else { // 在左边树上找
			rank = this.rank(node.leftNode, key)
		}
	} else {
		rank = 0
	}
	return
}

// Rank
func (this *BST) Rank(key *Comparable) (rank int) {
	return this.rank(this.root, key)
}

func (this *BST) deleteMin(node *BSTNode) (repNode *BSTNode) {
	if nil != node {
		if nil == node.leftNode {
			repNode = node.rightNode
		} else {
			node.leftNode = this.deleteMin(node.leftNode)
			node.size = node.leftNode.Size() + node.rightNode.Size() + 1
			repNode = node
		}
	}

	return
}

// DeleteMin 删除最小值
func (this *BST) DeleteMin() (err error) {
	if this.root == nil {
		err = errors.New("BST is Empty")
	} else {
		this.root = this.deleteMin(this.root)
	}
	return
}

func (this *BST) deleteMax(node *BSTNode) (repNode *BSTNode ) {
	if nil == node.rightNode {
		repNode = node.leftNode
	} else {
		node.rightNode = this.deleteMax(node.rightNode)
		node.size = node.leftNode.Size() + node.rightNode.Size() + 1
		repNode = node
	}

	return
}

// DeleteMin 删除最小值 FIXME 实现有误
func (this *BST) DeleteMax() (err error) {
	if this.root == nil {
		err = errors.New("BST is Empty")
	} else {
		this.root = this.deleteMax(this.root)
	}
	return
}

// 两个key之间的键的数量 [lo, hi] 两端都包含
func (this *BST) SizeBetween(lo *Comparable, hi *Comparable) (size int) {
	size = 0

	if lo.CompareTo(*hi) < 0 {
		loRank := this.Rank(lo)
		hiRank := this.Rank(hi) // 如果 hi 在 st内，那么 hi - lo + 1

		if this.Contains(hi) {
			size = hiRank - loRank + 1
		} else {
			size = hiRank - loRank
		}
	}

	return
}

// 检查每个树上的，然后返回。不过效率太低了。
// ( ∞, lo)  1
// [lo, hi]  2
// (hi, ∞ )  3
func (this *BST) keysBetween(node *BSTNode, keys *[]*Comparable, lo *Comparable, hi *Comparable) {

	// walk tree 中序遍历
	// 设置一个 status , 决定遍历区间
	if node != nil {
		compLo := node.key.CompareTo(*lo)
		compHi := node.key.CompareTo(*hi)

		if (compLo > 0) {
			this.keysBetween(node.leftNode, keys, lo, hi)
		}
		if (compLo >= 0 && compHi <= 0) {
			*keys = append(*keys, node.key)
		}
		if (compHi < 0) {
			this.keysBetween(node.rightNode, keys, lo, hi)
		}
	}

	return
}

// 返回 [lo, hi] 之间的键的列表
func (this *BST) KeysBetween(lo *Comparable, hi *Comparable) (keys []*Comparable) {
	keys = []*Comparable{}
	this.keysBetween(this.root, &keys, lo, hi)
	return
}

func readAndCountByBST(filename string, lengthThreshold int) (totalWordCount int, differendWordCount int) {
	return ReadAndCount(
		NewBST(),
		filename,
		lengthThreshold,
	)
}
