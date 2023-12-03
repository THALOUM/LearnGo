package main

import "fmt"

func main() {
	// 013 - Anonymous func:
	fmt.Println("Hello, 013 - Anonymous func:")
	fmt.Println("------------------------------------------------")

	// a named function with identifier:
	// func (r receiver) identifier (p parameter/s) (r return/s) { code }

	// an anonymous function:
	// func (p parameter/s) (could have r return/s) { code } (execution parentheses that will have the arguments)

	foo()

	// an anonymous function:
	func() {
		fmt.Println("Anonymous func ran.")
	}()

	// an anonymous function with a parameter:
	func(s string) {
		fmt.Println("Hello,", s)
	}("Amr!")

	// You can also assign an anonymous function to a variable and then invoke it using the variable name:
	multiply := func(a, b int) int {
		return a * b
	}

	fmt.Println(multiply(3, 5))
}

func foo() {
	fmt.Println("Foo ran.")
}
