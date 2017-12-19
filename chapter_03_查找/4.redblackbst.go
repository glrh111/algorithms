package main

import (
	"errors"
)

// 二叉树实现的 ST

/*
    二叉树的结点
 */
type RedBlackBSTNode struct {
	key *Comparable
	value interface{}
	leftNode, rightNode *BSTNode
	size int
}

func NewRedBlackBSTNode(key *Comparable, value interface{}, size int) *BSTNode {
	return &BSTNode{
		key: key,
		value: value,
		size: size,
	}
}

func (this *RedBlackBSTNode) Size() (size int) {
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
type RedBlackBST struct {
	root *BSTNode
}

func NewRedBlackBST() *RedBlackBST {
	return &RedBlackBST{
		root: nil,
	}
}


