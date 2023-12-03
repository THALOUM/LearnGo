package main

import "fmt"

func main() {
	// Exercise#70 - func return:
	fmt.Println("Hello, Exercise#70 - func return:")
	fmt.Println("------------------------------------------------")

	y := foo()
	fmt.Println(y())

	fmt.Printf("%T\n", foo)
	fmt.Printf("%T\n", y)

}

func foo() func() int {
	return func() int {
		return 42
	}
}
