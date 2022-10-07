package main

import "fmt"

func main(){
	var name string
	name = "Eko Kurniawan"
	fmt.Println(name)

	name = "Eko Khennedy"
	fmt.Println(name)
	var age = 30
	fmt.Println(age)

	/*
	decalar multiple variable

	var (
		firstName string
		lastName string
	)
	firstName = "Eko"
	lastName = "Khennedy"
	*/
	 var (
		firstName = "Eko"
		lastName = "Kurniawan"
	)
	fmt.Println(firstName, lastName)
	//asign value directly to variable without 'var' key using :=
	herName := "Lastari"

	fmt.Println(herName)

	herName = "Aidahlia"
	fmt.Println(herName)

}
