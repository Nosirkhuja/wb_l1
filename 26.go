package main

import (
	"fmt"
	"strings"
)

func isUniqueChars(str string) bool {
	str = strings.ToLower(str)
	var char_set [256]bool
	for i := 0; i < len(str); i++ {
		if char_set[str[i]] {
			return false
		}
		char_set[str[i]] = true
	}
	return true
}
func main() {
	s := "aAdsQ"
	d := "abcd"
	x := "самурай"
	v := "@!#$~`"
	fmt.Println(s, " :", isUniqueChars(s))
	fmt.Println(d, " :", isUniqueChars(d))
	fmt.Println(x, " :", isUniqueChars(x))
	fmt.Println(v, " :", isUniqueChars(v))
}
