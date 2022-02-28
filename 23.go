package main

import "fmt"

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func remove1(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	x := []int{4, 5, 6, 7, 88}
	fmt.Println(x)
	remove(x, 1)
	fmt.Println(x)
	fmt.Println(remove1(x, 1))
}
