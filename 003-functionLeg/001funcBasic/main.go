package main

import (
	"fmt"
)

// 006 - Methods:
type person struct {
    first string // Person's first name.
    last string // Person's last name.
    age int    // Person's age.
}

func main() {
    fmt.Println("Hello, Functions:-------")

    xi := []int {3,5,7,8,6,9}

    // calling a function:
    foo()
    bar("Amr")
    fmt.Println(aloha("Nerd"))
    s1 , n := dogYears("Amr", 4)
    fmt.Println(s1, n)
    fmt.Println("-----------------------------------------------")

    // 003 - Variadic parameter:
    fmt.Println("Sum of the input numbers:", sum())
    fmt.Println("Sum of the input numbers:", sum(1,2,3,4,5,6,7,8))
    fmt.Println("-----------------------------------------------")

    // 004 - Unfurling a slice:
    fmt.Println("Sum of the unfurling a slice:", sum(xi...))
    fmt.Println("-----------------------------------------------")

    // 006 - Methods: 
    amr := person {
        first: "Amr", 
        last: "Hassan",
        age: 28,
    }
    amr.speak()
    fmt.Println("-----------------------------------------------")

    // 005 - Defer:
    /*
        A "defer" statement invokes a function whose execution is deferred to the moment the surrounding 
        function returns, either because the surrounding function executed a return statement, reached 
        the end of its function body, or because the corresponding goroutine is panicking.

            DeferStmt = "defer" Expression .
        The expression must be a function or method call; it cannot be parenthesized. Calls of built-in 
        functions are restricted as for expression statements. 
    */
    fmt.Println("Defer statements: ")
    defer fmt.Println("world")
    fmt.Println("hello")
    fmt.Println("-----------------------------------------------")
}   

/*
    Define a function:
        func (r receiver) identifier (p parameter/s) (return/s) {code...}
*/
// no parameter, no returns:
func foo() {
    fmt.Println("I'm from foo.")
}

// parameters, no returns:
func bar(s string) {
    fmt.Println("My name is", s)
}

// parameters, 1 return:
func aloha(s string) string {
    return fmt.Sprint("They call me Mr.", s)
}

// 2 parameters, 2 returns:
func dogYears(name string, age int) (string, int) {
    age *= 7
    return fmt.Sprint(name, " is this old in dog years:"), age
}
/*--------------------------------------------------------------------------------------------------*/

// 003 - Variadic parameter: ...TYPE
// The final incoming parameter in a function signature.
//A function with variadic may be invoked with zero or more arguments for that parameter. 
/*
    Passing arguments to ... parameters
    If f is variadic with a final parameter p of type ...T, then within f the type of p is equivalent
    to type []T. If f is invoked with no actual arguments for p, the value passed to p is nil. 
    Otherwise, the value passed is a new slice of type []T with a new underlying array 
    whose successive elements are the actual arguments, which all must be assignable to T. 
    The length and capacity of the slice is therefore the number of arguments bound to p and may 
    differ for each call site.

    Given the function and calls
        func Greeting(prefix string, who ...string)
        Greeting("nobody")
        Greeting("hello:", "Joe", "Anna", "Eileen")
    within Greeting, who will have the value nil in the first call, and []string{"Joe", "Anna", "Eileen"} 
    in the second.

    If the final argument is assignable to a slice type []T and is followed by ..., it is passed unchanged 
    as the value for a ...T parameter. In this case no new slice is created.

    Given the slice s and call
        s := []string{"James", "Jasmine"}
        Greeting("goodbye:", s...)
    within Greeting, who will have the same value as s with the same underlying array
*/

// sum calculates the sum of a variadic list of integers(nums is a slice of int []int).
// It takes in any number of integers as input and returns their sum.
func sum(nums ...int) int {
    var total int = 0
    for _, num := range nums {
        total += num
    }
    return total
}
/*--------------------------------------------------------------------------------------------------*/

// 006 - Methods: in the function receiver:
func (p person) speak() {
    fmt.Printf("I'm %v %v, and I %v years old.\n", p.first, p.last, p.age)
}

