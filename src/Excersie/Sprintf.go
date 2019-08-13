// From the ex2.go code assing values to the identifiers declared and print them using Sprintf and assing the return value to the indenfier named s
// assign 42 to x , James Bond to y and true to z

package main

import "fmt"

// Decalaring variables at package level scope
var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}
