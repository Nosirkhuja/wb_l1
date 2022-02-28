package main

import (
	"fmt"
)

func TmGroup(tmp []float64) map[int][]float64 {
	groups := make(map[int][]float64)

	for _, val := range tmp {
		key := int(val/10) * 10
		groups[key] = append(groups[key], val)
	}
	return groups
}

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := TmGroup(temperatures)

	fmt.Printf("resul of splitting by groups %v\n", groups)
}
