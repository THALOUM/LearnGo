package main

import (
	"bytes"
	"fmt"
)

func main() {
	// 011 - Writer interface & writing to a byte buffer:
	fmt.Println("Hello, 011 - Writer interface & writing to a byte buffer:")
	fmt.Println("------------------------------------------------")

	b := bytes.NewBufferString("Hello, ")
	b.WriteString("Amr!.")
	fmt.Println(b.String())

	b.Reset()
	b.Write([]byte("Go is good."))
	fmt.Println(b.String())

}
