package main

import "fmt"

func main() {
	// Exercise#71 - callback:
	fmt.Println("Hello, Exercise#71 - callback:")
	fmt.Println("------------------------------------------------")

	fmt.Println(PrintSquare(5, square))
}

func square(n int) int {
	return n * n
}

func PrintSquare(a int, f func(int) int) string {
	square := f(a)
	return fmt.Sprintf("The square of %v is: %v.\n", a, square)
}
