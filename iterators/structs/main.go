package main

import "fmt"

type person struct{
	first string // firstName.
	last string // lastName.
	age int
}

type secretAgent struct {
	person
	ltk bool // Has a license to kill?
	first string // as the agent spy name.
}

func main() {
	amr := person{
		first: "Amr",
		last: "Hassan",
		age:28,
	}
	taha := person{
		first: "Taha",
		last: "Ahmed",
		age:21,
	}

	fmt.Println(amr)
	fmt.Printf("%T\t%#v\n", amr, amr)
	fmt.Println(taha)

	// we can access Fields values:
	fmt.Println(amr.age)

	// Embedded structs:
	bakr := secretAgent {
		person: person{
			first: "Bakr",
			last: "Hassan",
			age: 28,
		},
		first: "Bond", 
		ltk: true,
	}

	fmt.Printf("%#v\n", bakr)
	fmt.Println(bakr.ltk)
	fmt.Println(bakr.first)
	fmt.Println(bakr.person.first)

	// Anonymous structs:
	// useful if you just need struct in one place and you don't want to create an entire type.
	ahmed := struct{
		first string // firstName.
		last string // lastName.
		age int
	}{
		first: "Ahmed",
		last: "Sayed",
		age:29,
	}
	
	fmt.Printf("amr struct type: %T\n", amr)
	fmt.Printf("ahmed struct type: %T\n", ahmed)
	
}