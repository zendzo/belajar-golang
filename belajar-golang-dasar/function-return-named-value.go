package main

import "fmt"

func getfullName2() (firstName string, middleName string, lastName string) {
	// why we don't use := to initiate variable?
	// return value declared on function
	firstName = "Eko"
	middleName = "Kurinawan"
	lastName = "Khennedy"

	return
}

func main() {
	a, b, c := getfullName2()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
