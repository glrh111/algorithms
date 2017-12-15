package main

import (
	"strings"
	"os"
	"fmt"
	"bufio"
	"io"
)

type ComparableInterface interface {
	CompareTo(b Comparable) int
	Value() interface{}
}

type Comparable struct {
	value interface{}
}

func NewComparable(value interface{}) *Comparable {
	return &Comparable{value}
}

func (this *Comparable) CompareTo(b Comparable) (re int) {
	switch this.value.(type) {
	case int:
		thisValue := this.value.(int)
		thatValue := b.value.(int)
		if (thisValue > thatValue) {
			re = 1
		} else if (thisValue == thatValue) {
			re = 0
		} else {
			re = -1
		}
	case string:
		thisValue := this.value.(string)
		thatValue := b.value.(string)
		re = strings.Compare(thisValue, thatValue)
	default:
		panic("CompareTo only support int and string")
	}
	return
}

// 返回key的值
func (this *Comparable) Value() interface{} {
	return this.value
}

/*
   下面是ST的 interface
 */
 type SymbolTableInterface interface {
	// 更新操作
	Put(key *Comparable, value interface{})
	// 查找操作
	Get(key *Comparable) interface{}
	// 删除键值対
	Delete(key *Comparable)
	// 键值対个数
	Size() int
	// 是否包含某个值
	Contains(key *Comparable) bool
	// 是否为空
	IsEmpty() bool
	// 所有键的集合
	Keys() []*Comparable
}

type SortedSymbolTableInterface interface {
	SymbolTableInterface
	// 最小的键
	Min() (*Comparable, error)
	// 最大的键
	Max() (*Comparable, error)
	// 小于等于key的最大键
	Floor(*Comparable) (*Comparable, error)
	// 大于等于key的最小键
	Ceiling(*Comparable) (*Comparable, error)
	// 排名为k的键
	SelectKey(index int) (*Comparable, error)
	// 键k的排名
	Rank(key *Comparable) int
	// 删除最小值 表为空时返回错误
	DeleteMin() error
	// 删除最大值
	DeleteMax() error
	// lo, hi 之间键的数量
	SizeBetween(lo *Comparable, hi *Comparable) int
	// lo, hi 之间的所有键
	KeysBetween(lo *Comparable, hi *Comparable) []*Comparable
}

// 这个函数用于测试读取示例数据速度
func ReadAndCount(counter SymbolTableInterface, filename string, lengthThreshold int) (totalWordCount int, differendWordCount int) {
	totalWordCount, differendWordCount = 0, 0
	inputFile, inputError := os.Open(filename)
	if inputError != nil {
		fmt.Println("Open file error: ", inputError.Error())
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)

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
				everCount := counter.Get(NewComparable(word))
				if everCount == nil {
					everCount = 1
				} else {
					everCount = everCount.(int) + 1
				}
				counter.Put(NewComparable(word), everCount)
			}
		}
		if readError == io.EOF {
			break
		}
	}
	differendWordCount = counter.Size()
	return
}

