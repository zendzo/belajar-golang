package main

import "fmt"

func main() {
	/*
		const firstName string = "Eko"
		const lastName string = "Kurniawan"
		const age int = 766
	*/
	//multiple const declaration
	const (
		firstName string = "Eko"
		lastName  string = "Kurniawan"
		age       int    = 766
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
	// no error for unused const variable
	//fmt.Println(age)
}
