package main

import (
	"math/big"
	"fmt"
)

/*
   代表两个大数的 string 相乘
 */
func multiply(num1 string, num2 string) string {
	var (
		big1, big2 = big.NewInt(0), big.NewInt(0)
	)
	big1.SetString(num1, 10)
	big2.SetString(num2, 10)
	big1.Mul(big1, big2)
	return fmt.Sprintf("%v", big1)
}

func main() {
	fmt.Println(multiply("111111111111000000000000000000000000000000000000000000011", "111"))
}