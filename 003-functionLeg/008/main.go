package main

import (
	"fmt"
	"strconv"
)

type book struct {
	title string // Book's title.
}

func (b book) String() string {
	return fmt.Sprint("The title of the book is: ", b.title)
}

type count int 

func (c count) String() string {
	return fmt.Sprint("This is the number: ", strconv.Itoa(int(c)))
}

func main() {
	// 008 - Exploring the stringer interface:
	fmt.Println("Hello, 008 Exploring the stringer interface: ")
	/*
		type Stringer: 
		type Stringer interface {
			String() string
		}
	*/
	// so book and count are a type Stringer that allow you to print them differently.

	b := book {
		title: "The Great Gatsby.",
	}

	var n count = 42

	fmt.Println(b)
	fmt.Println(n)
}
