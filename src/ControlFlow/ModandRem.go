// Print modulus and reminder from 10 to 100 when divide by 4

package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Println(i, "i MOD: ", i%4, " and REM: ", i/4)
	}
}
