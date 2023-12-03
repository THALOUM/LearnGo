package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	first string // Person's first name.
}

func (p person) writeOut(w io.Writer) {
	w.Write([]byte(p.first))
}

func main() {
	// 012 - Writing to either a file or a byte buffer:
	fmt.Println("Hello, ")
	fmt.Println("------------------------------------------------")
	amr := person{
		first: "Amr",
	}

	// Creating a file:
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("Error: %s", err)
	} else {
		fmt.Printf("File Created Successfully in: ./%v\n", f.Name())
	}
	defer f.Close()

	var b bytes.Buffer

	amr.writeOut(f)

	amr.writeOut(&b) // & to point to a buffer.
	fmt.Println(b.String())
}
