package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 010 - Writer interface & writing to a file:
	fmt.Println("Hello, 010 - Writer interface & writing to a file:")
	fmt.Println("------------------------------------------------")

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("Error %s", err)
	}
	defer f.Close()

	s := []byte("Hello, Amr!")

	_, err = f.Write(s) // resigned err.
	if err != nil {
		log.Fatalf("Error %s", err)
	} else {
		fmt.Println("File created successfully.")
	}
}
