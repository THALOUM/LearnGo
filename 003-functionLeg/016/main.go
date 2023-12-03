package main

import "fmt"

func main() {
	// 016 - Callback:
	fmt.Println("Hello, 016 - Callback:")
	fmt.Println("------------------------------------------------")

	fmt.Printf("2 + 5 = %v\n", doOperation(2, 5, add))
	fmt.Printf("5 - 3 = %v\n", doOperation(5, 3, subtract))
	fmt.Printf("2 * 5 = %v\n", doOperation(2, 5, multiply))
	fmt.Printf("10 / 5 = %v\n", doOperation(10, 5, divide))

	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", doOperation)

}

func doOperation(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	return a / b
}
