package main

import "fmt"

func binarysearch(arr []int, x int) int {
	i := 0
	j := len(arr)
	for i != j {
		var m = (i + j) / 2
		if x == arr[m] {
			return m
		}
		if x < arr[m] {
			j = m
		} else {
			i = m + 1
		}
	}
	return -1
}

func main() {
	items := []int{1, 2, 3, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(" index of 10:", binarysearch(items, 10))
	fmt.Println(" index of 5:", binarysearch(items, 5))
	fmt.Println(" index of 20:", binarysearch(items, 20))
}
