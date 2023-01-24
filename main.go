package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ToJadenCase("How can mirrors be real if our eyes aren't real"))
}
func ToJadenCase(str string) string {
	// arr := strings.Split(str, " ")
	// fmt.Println(arr)
	result := ""
	result += strings.ToUpper(string(str[0]))
	for i := 1; i < len(str); i++ {

		if string(str[i-1]) == " " {
			result += strings.ToUpper(string(str[i]))
		} else {
			result += string(str[i])
		}
	}
	// arr1 := strings.ToUpper(arr[0])
	// fmt.Println(arr1)
	return result
}
