package main

import (
	"fmt"
	"log"
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

func logInfo(s fmt.Stringer) {
	log.Println("Log from 009,", s.String())
}

func main() {
	// 009 - Expanding on the stringer interface - wrapper func for logging:
	fmt.Println("Hello, 009 - Expanding on the stringer interface - wrapper func for logging")
	fmt.Println("------------------------------------------------")

	b := book{
		title: "The Great Gatsby.",
	}

	var n count = 42

	fmt.Println(b)
	fmt.Println(n)
	fmt.Println("------------------------------------------------")

	logInfo(b)
	logInfo(n)
}
