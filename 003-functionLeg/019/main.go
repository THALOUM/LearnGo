package main

import "fmt"

func main() {
	// 018 - Function fundamentals:
	// 019 -  Recursion:
	fmt.Println("Hello, 019 -  Recursion:")
	fmt.Println("------------------------------------------------")

	fmt.Printf("Factorial = %v\n", factorial(4))
	fmt.Printf("Factorial with looping = %v\n", factLoop(4))

}

func factorial(n int) int {
	fmt.Printf("n = %v\n", n)
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func factLoop(n int) int {
	total := 1
	for n > 0 {
		total *= n
		n--
	}
	return total
}
