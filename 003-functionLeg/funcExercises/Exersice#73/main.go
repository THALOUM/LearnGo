package main

import (
	"fmt"
	"time"
)

func main() {
	// Exercise#73 - wrapper:
	fmt.Println("Hello, Exercise#73 - wrapper:")
	fmt.Println("------------------------------------------------")

	fmt.Printf("The factorial of %v is: %v.\n", 5, factorial(5))
	TimedFunction(5, factorial)
}

func factorial(n int) int {
	if n < 1 {
		return 1
	}
	return n * factorial(n-1)
}

func TimedFunction(a int, f func(int) int) {
	start := time.Now()
	f(a)
	elapsedTime := time.Since(start)
	fmt.Printf("Factorial function take: %v.\n", elapsedTime)
}
