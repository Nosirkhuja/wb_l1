package main

import (
	"fmt"
)

func doSort(items []int, fst int, lst int) {
	if fst >= lst {
		return
	}
	i := fst
	j := lst
	x := items[(fst+lst)/2]

	for i < j {
		for items[i] < x {
			i++
		}
		for items[j] > x {
			j--
		}
		if i <= j {
			var tmp = items[i]
			items[i] = items[j]
			items[j] = tmp
			i++
			j--
		}
	}
	doSort(items, fst, j)
	doSort(items, i, lst)
}

func quicksort(arr []int) []int {
	length := len(arr)
	items := make([]int, length)
	copy(items, arr)
	doSort(items, 0, length-1)
	return items
}

func main() {
	items := []int{4, 1, 5, 3, 2}
	sortItems := quicksort(items)
	fmt.Println(sortItems)
}