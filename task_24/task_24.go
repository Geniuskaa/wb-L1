package main

import (
	"fmt"
	"math"
)

func main() {
	a := NewPoint(2, 5)
	b := NewPoint(-4, 10)

	fmt.Println(a.distanceToPoint(b))
	fmt.Println(b.distanceToPoint(a))
}

type Point struct {
	x int
	y int
}

func NewPoint(x int, y int) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) distanceToPoint(po *Point) int {
	x := (po.x - p.x) * (po.x - p.x)
	y := (po.y - p.y) * (po.y - p.y)
	dist := math.Sqrt(float64(x + y))

	return int(dist)
}
