package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	var res []string
	setOfString := make(map[string]bool, len(arr))

	for _, val := range arr {
		setOfString[val] = true
	}

	for key := range setOfString {
		res = append(res, key)
	}

	fmt.Printf("Arr: %v\nSet: %v\n", arr, res)
}
