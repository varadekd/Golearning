// Create a variable y of underlying type of user create type x
// Add these to the previous code (ex4.go)

package main

import "fmt"

type varade int

var a varade
var b int

func main() {
	fmt.Println("Value for x: ", a)
	fmt.Printf("Type for variable x: %T\n", a)

	// assinging value to x
	a = 42
	fmt.Println("Value for x after using assingement operator: ", a)

	b = int(a)
	fmt.Println("Value for b converstion: ", b)
}
