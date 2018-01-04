package stringst

/*
     单词查找树
 */

type Node struct {
	value interface{} // 为nil 代表无值
	nextList []*Node  // r 大小的数组
}

func NewNode(value interface{}, r int) *Node {
	return &Node{
		value: value,
		nextList: make([]*Node, r),
	}
}

type TrieST struct {
	size int
	root *Node
	r int       // 单词表大小 128
}

func NewTrieST() *TrieST {
	r := 128
	return &TrieST{
		size:0,
		root:NewNode(nil, r),
		r:r,
	}
}

// 返回找到的 node
func (st *TrieST) get(node *Node, key string, d int) (retNode *Node) {
	if node != nil {
		if d >= len(key) {
			retNode = node
		} else {
			retNode = st.get(node.nextList[charAt(key, d)], key, d+1)
		}
	}
	return
}

// 返回值. 如果没找到，返回 nil
func (st *TrieST) Get(key string) (retValue interface{}) {
	retNode := st.get(st.root, key, 0)
	if retNode != nil {
		retValue = retNode.value
	}
	return
}

func (st *TrieST) put(node *Node, key string, d int, value interface{}) (ifAdd bool) {
	charIndex := charAt(key, d)
	keyNode := node.nextList[charIndex]
	if keyNode == nil { // 新建结点
		node.nextList[charIndex] = NewNode(nil, st.r)
		ifAdd = st.put(node, key, d, value)
	} else { //
		if d >= len(key)-1 { // 检查完毕了
			if keyNode.value == nil {
				ifAdd = true
				st.size++
			}
			keyNode.value = value // 命中 ，更新值
		} else {
			ifAdd = st.put(keyNode, key, d+1, value)
		}
	}
	return
}

// PUT
func (st *TrieST) Put(key string, value interface{}) (ifAdd bool) {
	ifAdd = st.put(st.root, key, 0, value)
	return
}

func (st *TrieST) Size() int {
	return st.size
}

func (st *TrieST) contains(node *Node, key string, d int) (contains bool) {
	keyNode := node.nextList[charAt(key, d)]
	if keyNode != nil {
		if d >= len(key)-1 { // 命中，检查值
			contains = keyNode.value != nil
		} else {
			contains = st.contains(keyNode, key, d+1)
		}
	}
	return
}

func (st *TrieST) Contains(key string) (contains bool) {
	return st.contains(st.root, key, 0)
}

// collect 收集
func (st *TrieST) collect(node *Node, prefix string, c chan string, topCall bool) {
	if nil == node {
		return
	}
	if node.value != nil {
		c <- prefix
	}
	for i:=0; i<st.r; i++ {
		st.collect(node.nextList[i], prefix+indexToChar(i), c, false)
	}
	if topCall {
		close(c)
	}
}

func (st *TrieST) Keys() (c chan string) {
	c = make(chan string)
	go st.collect(st.root, "", c, true)
	return
}

func (st *TrieST) KeysWithPrefix(prefix string) (c chan string) {
	retNode := st.get(st.root, prefix, 0)
	c = make(chan string)
	if retNode != nil {
		go st.collect(retNode, prefix, c, true)
	} else {
		close(c)
	}
	return
}

func (st *TrieST) search(node *Node, s string, d int, length int) (l int) {
	if node == nil {
		return length
	}
	if len(s)-1 == d {     // s 查完了
		return d
	}
	if node.value != nil { // 继续往下走
		length = d
	}
	l = st.search(node.nextList[charAt(s, d)], s, d+1, length)
	return
}

// 返回值是key，它与 s 有最长的契合前缀
func (st *TrieST) LongestPrefixOf(s string) (key string) {
	length := st.search(st.root, s, 0, 0)
	return s[:length]
}

// FIXME 这里可以重构一下，在collect上加一个参数
func (st *TrieST) KeysThatMatch(pattern string) (c chan string) {
	c = make(chan string)

	go func() {
		for key := range st.Keys() {
			if IfMatch(pattern, key) {
				c <- key
			}
		}
		close(c)
	}()

	return
}

// 递归删除
func (st *TrieST) delete(node *Node, key string, d int) (n *Node, ifDelete bool) {
	if node == nil {
		return node, false
	}

	if len(key) == d {                 // 找到最后了
		if node.value != nil {         // 找到了这个值
			node.value = nil
			st.size--
			ifDelete = true
		}
	} else {
		node.nextList[charAt(key, d)], ifDelete = st.delete(node.nextList[charAt(key, d)], key, d+1) // 得检查一遍当前 node 否是为空
	}

	// 检查当前 node 是否为空
	nodeIfEmpty := true
	for i := 0; i < st.r; i++ {
		if node.nextList[i] != nil {
			nodeIfEmpty = false
			break
		}
	}
	if nodeIfEmpty { // 上面没有其他链接
		n = nil
	} else {
		n = node
	}

	return
}

// delete
func (st *TrieST) Delete(key string) (ifDelete bool) {
	st.root, ifDelete = st.delete(st.root, key, 0)
	return
}

// 返回字符对应的ascii码
func charAt(s string, d int) (ascii int) {
	return ASCII([]rune(s)[d])
}

func indexToChar(d int) (s string) {
	s = string(d)
	return
}

// ASCII of rune
func ASCII(r rune) (retr int) {
	retr = int(r)
	return
}

func IfMatch(key1 string, key2 string) (ifMatch bool) {
	if len(key1) == len(key2) {
		ifMatch = true
		for i:=0; i<len(key1); i++ {
			if !(key1[i] == '.' || key1[i] == key2[i]) {
				ifMatch = false
				break
			}
		}
	}
	return
}