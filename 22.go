package main

import (
	"fmt"
	"math/big"
)

func add(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func sub(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func multiply(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func divide(a, b *big.Int) *big.Int {
	return new(big.Int).Div(a, b)
}

func main() {

	bigIntA := big.NewInt(int64(1 << 34))
	bigIntB := big.NewInt(int64(1 << 28))

	fmt.Println(bigIntA)
	fmt.Println(bigIntB)

	fmt.Println("Adding is equal: ", add(bigIntA, bigIntB))

	fmt.Println("Subtracting is equal: ", sub(bigIntA, bigIntB))

	fmt.Println("Multiplying is equal: ", multiply(bigIntA, bigIntB))

	fmt.Println("Devision is equal: ", divide(bigIntA, bigIntB))
}
