package main

import "fmt"

/*
   test 一个 Sudoku 是否 valid
   http://sudoku.com.au/TheRules.aspx
   只能含有 [1-9] 9个元素
   + 每行的九个元素，不能重复
   + 每列的九个元素，也不能重复
   + 每个小单元格（一共9个单元格），不能重复
 */
func isValidSudoku(board [][]byte) bool {
	var (
		aux = []byte{}
	)
	// 检查横向
	for line:=0; line<9; line++ {
		aux = board[line]
		if hasRepeat(aux) {
			return false
		}
	}

	// 检查竖向
	for row:=0; row<9; row++ {
		aux = []byte{}
		for line:=0; line<9; line++ {
			aux = append(aux, board[line][row])
		}
		if hasRepeat(aux) {
			return false
		}
	}

	// 检查单元格
	fmt.Println("In 检查单元格：")
	for i:=0; i<9; i++ {
		line, row := i/3, i%3
		line2, row2 := 3*line, 3*row
		aux = []byte{}
		for line3:=line2; line3<line2+3; line3++ {
			for row3:=row2; row3<row2+3; row3++ {
				aux = append(aux, board[line3][row3])
			}
		}
		if hasRepeat(aux) {
			return false
		}
	}

	return true
}

func hasRepeat(b []byte) bool { // 检查 1 ~ 9 = 49 ~ 57
	aux := make([]bool, 256)
	for _, c := range b {
		switch {
		case 49 <= c && c <= 57: // 查看有无重复
			if aux[c] {
				return true
			}
			aux[c] = true
		case c == 46: // . 号
			continue
		default:
			return true
		}
	}
	return false
}

func main() {
	//b := [][]byte{['.','8','7','6','5','4','3','2','1'],['2','.','.','.','.','.','.','.','.'],['3','.','.','.','.','.','.','.','.'],['4','.','.','.','.','.','.','.','.'],['5','.','.','.','.','.','.','.','.'],['6','.','.','.','.','.','.','.','.'],['7','.','.','.','.','.','.','.','.'],['8','.','.','.','.','.','.','.','.'],['9','.','.','.','.','.','.','.','.']}
	b := []byte{46, 56, 55, 54, 53, 52, 51, 50, 49,}
}