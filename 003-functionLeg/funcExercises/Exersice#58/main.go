package main

import "fmt"

func main() {
	// Exercise#58:
	fmt.Println("Hello, Exercise#58:")
	fmt.Println("------------------------------------------------")

	x := foo()
	fmt.Println(x)

	s, y := bar()
	fmt.Println(s, y)

}

func foo() int {
	return 42
}

func bar() (string, int) {
	return "Hello", 43
}
