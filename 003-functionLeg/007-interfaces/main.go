package main

import "fmt"

type person struct {
	first string // Person's first name.
}

type secretAgent struct {
	person
	ltk bool // Is the agent has license to kill?
}

// interface:
type human interface {
	speak() // so any value of type has this method is a type human.
}

func main() {
	// 007 Interfaces & polymorphism:
	sa1 := secretAgent {
		person: person {
			first: "Amr",
		},
		ltk: true,
	}
	p2 := person {
		first: "jenny",
	}

	sa1.speak()
	p2.speak()
	fmt.Println("----------------------------------")

	saySomething(sa1)
	saySomething(p2)
}

func (p person) speak() {
	fmt.Printf("My name is %v.\n", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("I'm secret agent", sa.first)
}

func saySomething(h human) {
	h.speak()
}