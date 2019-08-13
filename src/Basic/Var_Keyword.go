// Declare vairable using var keyword at package level scope and do not assign the values to it
// Identifiers are x of type int, y of type string and z is of type boolean
// Print these values, and what is the value printed by the compilers are called

package main

import "fmt"

// Decalaring variables at package level scope

var x int
var y string
var z bool

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// The values printed by the compliers are called as Zero Value
}
