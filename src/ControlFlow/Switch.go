// Create a program that uses switch statement

package main

import "fmt"

func main() {
	a := "Hie"

	switch a {
	case "hi":
		fmt.Println("Print hi")
	case "hey":
		fmt.Println("Print hey")
	default:
		fmt.Println("Default a: ", a)
	}
}
