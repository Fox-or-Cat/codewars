package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(MaxRot(56789))
}
func MaxRot(n int64) int64 {
	str := strconv.FormatInt(n, 10)
	max := n
	for i := 0; i < len(str)-1; i++ {
		str = str[:i] + str[i+1:] + string(str[i])
		current, _ := strconv.ParseInt(str, 10, 64)
		if current > max {
			max = current
		}
	}
	return max
}
