package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func Distance(p1, p2 *Point) float64 {
	return math.Sqrt(math.Pow((p2.x-p1.x), 2) + math.Pow((p2.y-p1.y), 2))
}

func main() {
	point1 := NewPoint(1, 1)
	point2 := NewPoint(1, 2)

	fmt.Println("x1=", point1.x, " y1=", point1.y)
	fmt.Println("x2=", point2.x, " y2=", point2.y)
	fmt.Println("distance between two point is: ", Distance(point1, point2))
}
