package main

import "fmt"

type person struct {
	first string // Person's first name.
	age   int    // Person's age.
}

func (p person) speak() {
	fmt.Printf("My name is %v, I'm %v years old.\n", p.first, p.age)
}

func main() {
	// Exercise#61 - method:
	fmt.Println("Hello, Exercise#61 - method:")
	fmt.Println("------------------------------------------------")

	amr := person{
		first: "Amr",
		age:   28,
	}

	amr.speak()
}
