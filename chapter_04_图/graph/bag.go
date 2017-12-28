package graph

import (
	"strconv"
	"hash/fnv"
	//"fmt"
)

// 使用开放寻址hashST实现的BAG

type Key struct {
	value int
}

func NewKey(value int) *Key {
	return &Key{value}
}

func (k *Key) Value() int {
	return k.value
}

func (k *Key) IsEqual(anotherKey *Key) bool {
	return k.Value() == anotherKey.Value()
}

type Bag struct {
	keyList []*Key
	size int                // 实际存储的键值対的数量
	m int                   // 数组总大小
	minM int                // 设置的最小的数组的大小
}

func NewBag(initM int) *Bag {
	minM := 10
	if initM < minM {
		initM = minM
	}
	return &Bag{
		make([]*Key, initM),
		0,
		initM,
		minM,
	}
}

// resize cap 是新cap
func (b *Bag) resize(cap int) {
	that := NewBag(cap)
	for i := 0; i < b.m; i++ {
		thisKey := b.keyList[i]
		if thisKey != nil {
			that.Add(thisKey)
		}
	}
	*b = *that
}

func (b *Bag) hashIndex(key *Key) (index int) {
	repString := strconv.Itoa(key.Value())
	h := fnv.New32()
	h.Write([]byte(repString))
	index = int(h.Sum32() % uint32(b.m))
	return
}

func (b *Bag) nextIndex(nowIndex int) int {
	if nowIndex >= b.m - 1 {
		return 0
	} else {
		return nowIndex + 1
	}
}

func (b *Bag) Add(item *Key) {
	// 插入之前，判断 size
	if b.size > b.m / 2 {
		b.resize(2 * b.m)
	}

	// 线性查找法插入

	hashIndex := b.hashIndex(item)

	for {
		thisKey := b.keyList[hashIndex]

		if thisKey == nil { // 没值，设置新值
			b.keyList[hashIndex] = item
			b.size++
			break
		} else {

			if item.IsEqual(thisKey) { // 命中, 什么也不干
				break
			} else {                          // hashIndex + 1
				hashIndex = b.nextIndex(hashIndex)
			}
		}
	}
}

func (b *Bag) IsEmpty() bool {
	return 0 == b.Size()
}

func (b *Bag) Size() int {
	return b.size
}

func (b *Bag) Iterator() func () (int, bool) {
	nowSize := 0
	i := 0
	return func () (int, bool) {
		for {
			if !(i<b.m && nowSize<b.size) {
				break
			}

			thisKey := b.keyList[i]

			if thisKey == nil {
				i++
				continue
			}
			nowSize++
			i++
			return thisKey.Value(), true
		}
		return 0, false
	}
}

func (b *Bag) IteratorChan() chan int {
	nowSize := 0
	i := 0
	c := make(chan int)
	go func () {
		for {
			if !(i<b.m && nowSize<b.size) {
				break
			}

			thisKey := b.keyList[i]

			if thisKey == nil {
				i++
				continue
			}
			nowSize++
			i++
			c <- thisKey.Value()
		}
		close(c)
	}()
	return c
}

