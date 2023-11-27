package main

import "fmt"

type person struct{
	first string // Person's first name.
	last string // Person's last name.
	flavors []string // Person's favorite ice cream flavors.
}

type engine struct{
	electric bool
}

type vehicle struct{
	engine  
	make string
	model string
	doors int
	color string
}

func main() {
	// Exercise##53 - struct with slice:
	amr := person{
		first: "Amr",
		last: "Hassan",
		flavors: []string {"chocolate", "banana", "passion fruit with mango and guava"},
	}
	taha := person{
		first: "Taha",
		last: "Ahmed",
		flavors: []string {"Strawberry", "Chocolate"},
	}

	fmt.Printf("amr: %#v\n", amr)
	fmt.Printf("taha: %#v\n", taha)

	for index, flavor := range amr.flavors {
		fmt.Printf("amr's favorite flavor %v: %v\n", index+1, flavor)
	}

	for index, flavor := range taha.flavors {
		fmt.Printf("taha's favorite flavor %v: %v\n", index+1, flavor)
	}
	fmt.Println("------------------------------------------------")

	// Exercise #54 - map struct:
	people := make(map[string]person)
	people[amr.last] = amr
	people[taha.last] = taha
	/*
		people := map[string]person {
			amr.last: amr,
			taha.last: taha,
		}
	*/

	fmt.Println(people)

	for _,person := range people{
		fmt.Printf("--------%v's favorite flavor--------\n", person.first)
		for _,flavor := range person.flavors{
			fmt.Println(flavor)
		}
	}
	fmt.Println("------------------------------------------------")

	// Exercise #55 - embed struct:
	/*
	Car 1:
        Maker: Toyota
        Model: Camry
        Color: Silver
        Number of Doors: 4
        Electric Engine: No

    Car 2:
        Maker: Tesla
        Model: Model S
        Color: Red
        Number of Doors: 2
        Electric Engine: Yes
	*/
	car1 := vehicle{
		engine: engine{electric: false},
		make: "Toyota",
        model: "Camry",
        doors: 4,
        color: "Silver",
	}

	car2 := vehicle{
		engine: engine{electric: true},
		make: "Tesla",
        model: "Model S",
        doors: 2,
        color: "Red",
	}

	fmt.Printf("Car1: %#v\n", car1)
	fmt.Printf("Car2: %#v\n", car2)

	fmt.Printf("Is car1 has an electric engine? %v\n", car1.engine.electric)
	fmt.Printf("Is car2 has an electric engine? %v\n", car2.engine.electric)
	fmt.Println("------------------------------------------------")

	// Exercise #56 - anonymous struct:
	bakr := struct{
		first string // first name.
		friends []string // friends name.
		favoriteDrink string 
	}{
		first: "Bakr",
		friends: []string {"Amr", "Ahmed"},
		favoriteDrink: "Tea",
	}

	fmt.Printf("%#v\n", bakr)
	fmt.Println("------------------------------------------------")

}