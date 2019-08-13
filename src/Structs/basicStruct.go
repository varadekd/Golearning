// Create struct named person that stores first name, Last name and Fav ice creams

package main

import "fmt"

type person struct {
	firstName    string
	lastName     string
	favIceCreams []string
}

func main() {
	p1 := person{
		firstName:    "Kushagra",
		lastName:     "Varade",
		favIceCreams: []string{"Butterscotch", "Vanilla", "Kaju pista", "Malai"},
	}

	p2 := person{
		firstName:    "Yash",
		lastName:     "Pauranik",
		favIceCreams: []string{"Butterscotch", "Vanilla"},
	}

	fmt.Println("P1:", p1)
	fmt.Println("P2:", p2)
}
