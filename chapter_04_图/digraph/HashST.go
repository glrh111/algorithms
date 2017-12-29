package digraph

import (
	"strconv"
	"hash/fnv"
	"strings"
)

// 线性探测法 ST
// 使用并行数组，一个保存键，另一个保存值

// 如何动态扩展数组大小？
// 1/ Delete() 操作后，在 0 < (size/M) <= 1/8 resize(M / 2)
// 2/ Put()    操作厚，在 size/M <= 1/2       resize(M * 2)
// 3/ 思考，需要最小保证M为 10. 再小就没意义了。我觉得。

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

type LinearProbingHashST struct {
	keyList []*Comparable
	valueList []interface{} // 容量为m
	size int                // 实际存储的键值対的数量
	m int                   // 数组总大小
	minM int                // 设置的最小的数组的大小
}


// 构造一个新的hash表
func NewLinearProbingHashST(initM int) *LinearProbingHashST {
	minM := 10
	if initM < minM {
		initM = minM
	}
	return &LinearProbingHashST{
		make([]*Comparable, initM),
		make([]interface{}, initM),
		0,
		initM,
		minM,
	}
}

// resize cap 是新cap
func (st *LinearProbingHashST) resize(cap int) {
	that := NewLinearProbingHashST(cap)
	for i := 0; i < st.m; i++ {
		thisKey := st.keyList[i]
		if thisKey != nil {
			that.Put(thisKey, st.valueList[i])
		}
	}
	*st = *that
}

// hashCode
func (st *LinearProbingHashST) hashIndex(key *Comparable) (index int) {
	repString := ""
	switch key.Value().(type) {
	case int:
		repString = strconv.Itoa(key.Value().(int))
	case string:
		repString = key.Value().(string)
	default:
		panic("type not support")
	}
	h := fnv.New32()
	h.Write([]byte(repString))
	index = int(h.Sum32() % uint32(st.m))
	return
}

func (st *LinearProbingHashST) Size() int {
	return st.size
}

// nextIndex
func (st *LinearProbingHashST) nextIndex(nowIndex int) int {
	if nowIndex >= st.m - 1 {
		return 0
	} else {
		return nowIndex + 1
	}
}

// GET
func (st *LinearProbingHashST) Get(key *Comparable) (value interface{}) {

	if nil == key {
		return nil
	}

	hashIndex := st.hashIndex(key)

	for {
		thisKey := st.keyList[hashIndex]

		if thisKey  == nil { // 没值，未命中
			value = nil
			break
		} else {

			if thisKey.CompareTo(*key) == 0 { // 命中
				value = st.valueList[hashIndex]
				break
			} else {                          // hashIndex + 1
				hashIndex = st.nextIndex(hashIndex)
			}
		}
	}

	return value
}

// PUT
func (st *LinearProbingHashST) Put(key *Comparable, value interface{}) {
	// 插入之前，判断 size
	if st.size > st.m / 2 {
		st.resize(2 * st.m)
	}

	// 线性查找法插入

	hashIndex := st.hashIndex(key)

	for {
		thisKey := st.keyList[hashIndex]

		if thisKey  == nil { // 没值，设置新值
			st.keyList[hashIndex] = key
			st.valueList[hashIndex] = value
			st.size++
			break
		} else {

			if thisKey.CompareTo(*key) == 0 { // 命中, 更新
				st.valueList[hashIndex] = value
				break
			} else {                          // hashIndex + 1
				hashIndex = st.nextIndex(hashIndex)
			}
		}
	}

}

// Delete 找到之后，其后的所有连续元素，都得重新插入到散列表里边。
func (st *LinearProbingHashST) Delete(key *Comparable) {

	hashIndex := st.hashIndex(key)

	for {
		thisKey := st.keyList[hashIndex]

		if thisKey  == nil { // 未命中
			break
		} else {

			if thisKey.CompareTo(*key) == 0 { // 命中, 删除
				st.keyList[hashIndex] = nil
				st.valueList[hashIndex] = nil

				// 之后的键簇中的元素重新插入
				for {
					hashIndex = st.nextIndex(hashIndex)

					if st.keyList[hashIndex] == nil { // 其后没有元素
						break
					} else {
						st.Put(st.keyList[hashIndex], st.valueList[hashIndex])
					}

				}
				st.size--

				// 删除之后，resize
				if (st.m >= st.minM && st.size <= st.m / 8) {
					st.resize(st.m / 2)
				}

				break
			} else {                          // hashIndex + 1
				hashIndex = st.nextIndex(hashIndex)
			}
		}
	}


}

// Contains 是否包含某个元素
func (st *LinearProbingHashST) Contains(key *Comparable) bool {

	hashIndex := st.hashIndex(key)

	for {
		thisKey := st.keyList[hashIndex]

		if thisKey  == nil { // 没值，设置新值
			return false
		} else {

			if thisKey.CompareTo(*key) == 0 { // 命中, 更新
				return true
			} else {                          // hashIndex + 1
				hashIndex = st.nextIndex(hashIndex)
			}
		}
	}

}

func (st *LinearProbingHashST) IsEmpty() bool {
	return 0 == st.size
}

func (st *LinearProbingHashST) Keys() (keys []*Comparable) {
	keys = make([]*Comparable, st.Size())
	actualIndex := 0
	for i := 0; i < st.m; i++ {
		thisKey := st.keyList[i]
		if thisKey != nil {
			keys[actualIndex] = thisKey
			actualIndex ++
		}
	}

	return
}

//// 读取测试用例
//func readAndCountByLinearProbingHashST(filename string, lengthThreshold int) (totalWordCount int, differendWordCount int) {
//	return ReadAndCount(
//		NewLinearProbingHashST(10),
//		filename,
//		lengthThreshold,
//	)
//}

