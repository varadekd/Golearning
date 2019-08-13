// Create your own type of int
// Create new variable of user type "x" using var key word
// Print value for "x"
// Print type for "x"
//  Assign value 42 to "x"
// print out new value for x

package main

import "fmt"

type varade int

var a varade

func main() {
	fmt.Println("Value for x: ", a)
	fmt.Printf("Type for variable x: %T\n", a)

	// assinging value to x
	a = 42
	fmt.Println("Value for x after using assingement operator: ", a)

}
