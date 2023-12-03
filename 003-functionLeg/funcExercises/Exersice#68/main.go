package main

import "fmt"

func main() {
	// Exercise#68 - anonymous func:
	fmt.Println("Hello, Exercise#68 - anonymous func:")
	fmt.Println("------------------------------------------------")

	func(s string) {
		fmt.Println("Hello,", s)
	}("Amr!")

	x := func() int {
		return 42
	}

	fmt.Println(x())
}
