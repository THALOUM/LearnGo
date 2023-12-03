package main

import (
	"fmt"
	"math"
)

type SQUARE struct {
	length float64
	width  float64
}

type CIRCLE struct {
	radius float64
}

type SHAPE interface {
	area()
}

func (sq SQUARE) area() {
	area := sq.length * sq.width
	fmt.Printf("The area is: %v\n", area)
}

func (c CIRCLE) area() {
	area := math.Pi * math.Pow(c.radius, 2)
	fmt.Printf("The area is: %v\n", area)
}

func info(shape SHAPE) {
	shape.area()
}

func main() {
	// Exercise#62 - interfaces:
	fmt.Println("Hello, Exercise#62 - interfaces:")
	fmt.Println("------------------------------------------------")

	square := SQUARE{
		length: 4,
		width:  3,
	}
	info(square)

	circle := CIRCLE{
		radius: 4,
	}
	info(circle)
}
