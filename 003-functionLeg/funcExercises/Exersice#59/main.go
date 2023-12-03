package main

import "fmt"

func main() {
	// Exercise#59 - variadic func:
	fmt.Println("Hello, Exercise#59 - variadic func:")
	fmt.Println("------------------------------------------------")

	xi := []int{1, 4, 5, 6, 8}

	fmt.Printf("Sum of unfurling slice: %v\n", sum(xi...))
	fmt.Printf("Sum of the slice without unfurling: %v\n", barSum(xi))
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func barSum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
