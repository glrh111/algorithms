package main

import (
	"fmt"
	"bufio"
	"strings"
	"io"
	"os"
	"errors"
)

// 二分查找法实现的符号表
// 已知问题：深度copy问题。暂时不改。

type BinarySearchST struct {
	keys []*Comparable
	values []interface{}
	size int
}

func NewBinarySearchST() *BinarySearchST {
	return &BinarySearchST{
		keys: []*Comparable{},
		values: []interface{}{},
		size: 0,
	}
}

// 返回小于key的元素的数量
func (this *BinarySearchST) Rank(key *Comparable) (rank int) {
	lo, hi := 0, this.size-1
	rank = 0
	for ; lo<= hi; {
		mid := (lo + hi) / 2
		compRe := this.keys[mid].CompareTo(*key)
		if (compRe == 0) { // 刚好相等
			rank = mid
			return
		} else if (compRe < 0) {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	rank = lo
	return
}

// 取出元素
func (this *BinarySearchST) Get(key *Comparable) (value interface{}) {
	rank := this.Rank(key)
	value = nil
	if (rank < this.size && this.keys[rank].CompareTo(*key) == 0) {
		value = this.values[rank]
	}
	return
}

// 设置元素的值
func (this *BinarySearchST) Put(key *Comparable, value interface{}) {
	rank := this.Rank(key)
	// 找到了这个元素，那么更新 value
	if (rank < this.size && this.keys[rank].CompareTo(*key) == 0) {
		this.values[rank] = value
	} else { // 没有找到，在rank位置上插入一个元素. 防止越界
		if (rank>=this.size) { // 在末尾追加
			this.keys = append(this.keys, key)
			this.values = append(this.values, value)
		} else {               // 在中间追加
			this.keys = append(this.keys[:rank], append([]*Comparable{key}, this.keys[rank:]...)...)
			this.values = append(this.values[:rank], append([]interface{}{value}, this.values[rank:]...)...)
		}

		this.keys[rank] = key
		this.values[rank] = value
		this.size++
	}
}

// 删除一个元素: 找到元素的话，直接删除
func (this *BinarySearchST) Delete(key *Comparable) {
	rank := this.Rank(key)
	// 找到了这个元素，执行删除操作
	if (rank < this.size && this.keys[rank].CompareTo(*key) == 0) {
		if (rank == this.size - 1) {
			this.keys = this.keys[:rank]
			this.values = this.values[:rank]
		} else {
			this.keys = append(this.keys[:rank], this.keys[rank+1:]...)
			this.values = append(this.values[:rank], this.values[rank+1:]...)
		}
		this.size--
	}
}

// 是否包含某个元素
func (this *BinarySearchST) Contains(key *Comparable) (ifContains bool) {
	rank := this.Rank(key)
	ifContains = false
	// 找到了这个元素，执行删除操作
	if (rank < this.size && this.keys[rank].CompareTo(*key) == 0) {
		ifContains = true
	}
	return
}

// 是否为空
func (this *BinarySearchST) IsEmpty() bool {
	return this.size == 0
}

// size
func (this *BinarySearchST) Size() int {
	return this.size
}

// keys copy 一份出来
func (this *BinarySearchST) Keys() []*Comparable{
	b := make([]*Comparable, len(this.keys))
	copy(b, this.keys)
	return b
}

// min 最小的key
func (this *BinarySearchST) Min() (key *Comparable, err error) {
	if this.IsEmpty() {
		key = nil
		err = errors.New("ST is empty")
	} else {
		key = this.keys[0]
		err = nil
	}
	return
}

// max 最大的key
func (this *BinarySearchST) Max() (key *Comparable, err error) {
	if this.IsEmpty() {
		key = nil
		err = errors.New("ST is empty")
	} else {
		key = this.keys[this.size-1]
		err = nil
	}
	return
}

// Floor 小于等于key的最大键
func (this *BinarySearchST) Floor(key *Comparable) (reKey *Comparable, err error) {
	rank := this.Rank(key)
	if (rank < this.size && this.keys[rank].CompareTo(*key)==0) { // 复制找到这个key返回
		reKey = &Comparable{}
		*reKey = *key
	} else if (rank <= 0) {                                       // 返回一个错误
		err = errors.New("No return value")
	} else {                                                      // 复制一份前一个元素，返回
		reKey = &Comparable{}
		*reKey = *this.keys[rank-1]
	}
	return
}

// Ceiling 大于等于key的最大键
func (this *BinarySearchST) Ceiling(key *Comparable) (reKey *Comparable, err error) {
	rank := this.Rank(key)
	if (rank >= this.size) {                               // 返回一个错误
		err = errors.New("No return value")
	} else {                                                      // 复制一份前一个元素，返回
		reKey = &Comparable{}
		*reKey = *this.keys[rank]
	}
	return
}

// SelectKey 找出排名为某int 的键
func (this *BinarySearchST) SelectKey(rank int) (key *Comparable, err error) {
	if (0 <= rank && rank < this.size) {
		key = &Comparable{}
		*key = *this.keys[rank]
	} else {
		err = errors.New("Rank out of range")
	}
	return
}

// DeleteMin
func (this *BinarySearchST) DeleteMin() (err error) {
	if this.IsEmpty() {
		err = errors.New("ST is empty.")
	} else {
		this.Delete(this.keys[0])
	}
	return
}

// DeleteMax
func (this *BinarySearchST) DeleteMax() (err error) {
	if this.IsEmpty() {
		err = errors.New("ST is empty.")
	} else {
		this.Delete(this.keys[this.size-1])
	}
	return
}

// SizeBetween(lo, hi) [lo, hi] 之间键的数量
func (this *BinarySearchST) SizeBetween(lo *Comparable, hi *Comparable) (size int) {
	size = 0
	if lo.CompareTo(*hi) != 1 {
		loRank := this.Rank(lo)
		hiRank := this.Rank(hi)
		// 如果hiRank 这个位置的元素存在的话，hi-lo+1, 否则hi-lo
		if (hiRank < this.size && this.keys[hiRank].CompareTo(*hi) == 0) {
			size = hiRank - loRank + 1
		} else {
			size = hiRank - loRank
		}
	}

	return
}

// KeysBetween(lo, hi) [lo, hi] 之间的所有键
func (this *BinarySearchST) KeysBetween(lo *Comparable, hi *Comparable) (keys []*Comparable) {
	keys = []*Comparable{}
	if lo.CompareTo(*hi) != 1 {
		loRank := this.Rank(lo)
		hiRank := this.Rank(hi)
		// 如果hiRank 这个位置的元素存在的话，hi-lo+1, 否则hi-lo
		size := 0
		if (hiRank < this.size && this.keys[hiRank].CompareTo(*hi) == 0) {
			size = hiRank - loRank + 1
		} else {
			size = hiRank - loRank
		}
		// 判断lo是否存在，因为需要判断是否包含lo所在的元素
		hiIndex := loRank + size
		//fmt.Println("loRank, ", loRank, "hiRank, ", hiRank, "size, ", size, "hiIndex, ", hiIndex)
		if hiIndex >= loRank {
			// 需要分配空间
			keys = make([]*Comparable, size)
			copy(keys, this.keys[loRank:hiIndex])
		}

	}

	return
}

// 二分查找法实现的符号表的实际应用
func readAndCountByBinarySearchST(filename string, lengthThreshold int) (totalWordCount int, differendWordCount int) {
	totalWordCount, differendWordCount = 0, 0
	inputFile, inputError := os.Open(filename)
	if inputError != nil {
		fmt.Println("Open file error: ", inputError.Error())
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)

	// 构造一个ST
	st := NewBinarySearchST()

	for {
		inputString, readError := inputReader.ReadString('\n')
		// 去掉 \n
		inputString = strings.Trim(inputString, "\n")
		wordList := strings.Split(inputString, " ")
		for _, word := range wordList {
			//fmt.Println(word, " ", len(word))
			if len(word) >= lengthThreshold {
				totalWordCount += 1
				// 首先查找在不在
				everCount := st.Get(NewComparable(word))
				if everCount == nil {
					everCount = 1
				} else {
					everCount = everCount.(int) + 1
				}
				st.Put(NewComparable(word), everCount)
			}
		}
		if readError == io.EOF {
			break
		}
	}
	differendWordCount = st.Size()
	return
}
