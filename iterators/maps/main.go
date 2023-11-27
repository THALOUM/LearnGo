package main

import "fmt"

func main() {
	a := map[string]int {
		"Amr": 0001,
		"Taha": 0002,
		"Ahmed": 0003,
	}

	// add element to a map:
	a["Ezz"] = 004

	fmt.Println(a["Amr"])
	fmt.Printf("%#v\n", a)

	// We can create a map using make:
	b := make(map[string]int)
	b["Bakr"] = 005
	b["Tarek"] = 006

	fmt.Printf("%#v\n", b)

	// Map - for range over a map:
	for key, value := range a {
		fmt.Printf("%v's number is: %v\n", key, value)
	}

	// Map - delete element: by delete built in function:
	delete(a, "Ahmed")
	fmt.Printf("%#v\n", a)

	// accessing key that don't exist:
	fmt.Printf("accessing key that don't exist: \n")
	delete(a, "Ahmed") // won't panic.
	fmt.Println(a["Ahmed"]) // return zero because we declare the map values as int.
	// with that we can't know if the key is not exist or it's actual value zero.
	// to solve that we use comma ok idiom.
	fmt.Println("--------------------------")

	// Map - comma ok idiom: ok value is true or false according to the key exist or not:
	if value, ok := a["Ahmed"]; ok{ 
		fmt.Printf("The key exist and it's value: %v.\n", value)
	} else {
		fmt.Println("Key not found!")
	}

	// Map - counting words in a book:
	c := 0
	fmt.Println(c)
	c++
	fmt.Println(c)

	fmt.Println("-------")

	m := make(map[string]int)
	fmt.Println(m["beautiful"])
	fmt.Printf("%#v\n", m)
	m["beautiful"]++
	fmt.Printf("%#v\n", m)
	fmt.Println(m["beautiful"])
	m["beautiful"]++
	fmt.Println(m["beautiful"])
	m["beautiful"]++
	fmt.Println(m["beautiful"])

}