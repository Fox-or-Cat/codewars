package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsTriangle(6, 7, 8))
}

func IsTriangle(a, b, c int) bool {
	if a+b > c && a+c > b && b+c > a {
		return true
	}
	return false
}
