package main

import "fmt"

func main() {
	// Exercise#64 - tests in go #2 - unit tests:
	fmt.Println("Hello, Exercise#64 - tests in go #2 - unit tests:")
	fmt.Println("------------------------------------------------")

	x := DoMath(5, 5, Add)
	fmt.Println(x)

	y := DoMath(5, 2, Subtract)
	fmt.Println(y)
}

// Add takes two integers as input and returns their sum.
func Add(a, b int) int {
	return a + b
}

// Subtract takes two integers, a and b, and returns their difference.
func Subtract(a, b int) int {
	return a - b
}

// DoMath is a function that takes two integers, a and b, and a function f.
// It returns the result of applying the function f to the integers a and b.
func DoMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}
