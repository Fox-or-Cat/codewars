package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	fmt.Println(Calc("ifkhchlhfd"))
}
func Calc(s string) int {
	runeArray := []rune(s)
	str1 := ""
	for i := 0; i < len(runeArray); i++ {
		str1 += strconv.FormatInt(int64(runeArray[i]), 10)
	}
	str2 := ""
	for i := 0; i < len(str1); i++ {
		if string(str1[i]) == "7" {
			str2 += "1"
		} else {
			str2 += string(str1[i])
		}
	}
	total1, _ := new(big.Int).SetString(str1, 10)
	total2, _ := new(big.Int).SetString(str2, 10)
	diff := new(big.Int)
	diff.Sub(total1, total2)
	diffStr := diff.String()
	result := 0
	for i := 0; i < len(diffStr); i++ {
		num, _ := strconv.Atoi(string(diffStr[i]))
		result += num
	}
	return result
}
