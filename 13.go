package main

import "fmt"

func main() {
	//1 метод - математика и никакой магии
	a := 1
	b := 2
	fmt.Println(a, " ", b)
	a = a + b
	b = b - a
	b = -b
	a = a - b
	fmt.Println(a, " ", b)
	//2 подобно Питоновской a,b=b,a
	c := 3
	d := 4
	fmt.Println(c, " ", d)
	c, d = d, c
	fmt.Println(c, " ", d)
}
