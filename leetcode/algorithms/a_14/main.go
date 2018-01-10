package main

import "fmt"

type treeNode struct {
	value bool // 代表是否有单词
	nodeList []*treeNode
	byteList []byte // nodeList 的非空链接
}

func newNode() *treeNode {
	return &treeNode{false, make([]*treeNode, 256), []byte{}}
}

type tree struct {
	root *treeNode
}

func newTree(strs []string) (t *tree) {
	t = &tree{newNode()}
	for _, str := range strs {
		t.put(str)
	}
	return 
}

func (t *tree) reput(s string, node *treeNode, d int) *treeNode {

	if node.nodeList[s[d]] == nil { // 创建新的 node
		node.nodeList[s[d]] = newNode()
		node.byteList = append(node.byteList, s[d])
	}

	if d >= len(s) - 1 { // 更新里边的值
		node.nodeList[s[d]].value = true
	} else {
		node.nodeList[s[d]] = t.reput(s, node.nodeList[s[d]], d+1)
	}

	return node
}

func (t *tree) put(s string) {
	if s == "" {
		t.root.value = true
		return
	}
	t.root = t.reput(s, t.root, 0)
}

func (t *tree) longestprefix() (s string) {

	currentNode := t.root // 还得判断单词是否中止了。

	for {

		outLinkAmount := len(currentNode.byteList)
		if outLinkAmount >= 2 || outLinkAmount == 0 || currentNode.value { // 到尽头了，或者指向不只一个其他字母
			break
		}
		s += string(currentNode.byteList[0])
		currentNode = currentNode.nodeList[currentNode.byteList[0]]
	}
	return
}

/*
   找出一串单词的最长前缀。用单词查找树解决？Good。说干就干。

   FIXME 这种方法只超过了 3% 的人。FUCK
 */
func longestCommonPrefix(strs []string) string {
	return newTree(strs).longestprefix()
}

func main() {
	strs := []string{"wocao", "woooonidaye", "wawjfesf"}
	strs2 := []string{"b", "b"}
	fmt.Println(longestCommonPrefix(strs), longestCommonPrefix(strs2))
}
