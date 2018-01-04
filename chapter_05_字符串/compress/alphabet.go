package compress

import "math"

type Alphabet struct {
	r int // 基数
	charList []rune // 字母表中的字符列表
	charST  *LinearProbingHashST  // 一个符号表，char -> index
}

func NewAlphabet(s string) (ab *Alphabet) {
	r := []rune(s)

	ab = &Alphabet{
		r: len(r),
		charList: r,
		charST: NewLinearProbingHashST(10),
	}

	for index, sr := range r {
		ab.charST.Put(NewComparable(sr), index)
	}

	return
}

func (ab *Alphabet) ToChar(index int) (r rune) {
	r = ab.charList[index]
	return
}

func (ab *Alphabet) ToIndex(char rune) (index int) {
	index = ab.charST.Get(NewComparable(char)).(int)
	return
}

func (ab *Alphabet) Contains(char rune) (contains bool) {
	return ab.charST.Contains(NewComparable(char))
}

func (ab *Alphabet) R() int {
	return ab.r
}

// 字母表的 index 需要多少位才能表示？ log_(2)r
func (ab *Alphabet) LgR() int {
	return int(math.Ceil(math.Log2(float64(ab.r))))
}

func (ab *Alphabet) ToIndices(r []rune) (indexList []int) {
	indexList = make([]int, len(r))
	for index, sr := range r {
		indexList[index] = ab.ToIndex(sr)
	}
	return
}

func (ab *Alphabet) ToChars(indexList []int) (r []rune) {
	r = make([]rune, len(indexList))
	for i, index := range indexList {
		r[i] = ab.ToChar(index)
	}
	return
}

var (
	DNAALPHABET = NewAlphabet("ACTG") // 基因组字母表
	BINARYALPHABET = NewAlphabet("01") // 二进制
	LOWERCASEALPHABET = NewAlphabet("abcdefghijklmnopqrstuvwxyz")
)
