package main

import (
	"fmt"
	"math"
)

// both rect and circle has area so we can use that
// if we access this way only area is accessible
// not the radius, width or height
type shape interface {
	area() float32
}

type circle struct {
	radius float32
}

type rect struct {
	width  float32
	height float32
}

func (r rect) area() float32 {
	return r.width * r.height
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func getArea(s shape) float32 {
	return float32(s.area())
}

func main() {
	c1 := circle{3.0}
	r1 := rect{5, 8}
	shapes := []shape{c1, r1}

	for _, shape := range shapes {
		// fmt.Println(getArea(shape))
		fmt.Println(shape.area())
	}
}
