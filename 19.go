package main

import (
	"fmt"
)

func reverse(str string) string {
	runes := []rune(str)
	left, right := 0, len(runes)-1

	for left <= right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
	return string(runes)
}
func main() {
	s := "самурай"
	fmt.Println(s, " -> ", reverse(s))
}
