package main

import "fmt"

func getfullName() (string, string, string) {
	return "Eko", "Kurniawan", "Khennedy"
}

func main() {
	firstName, _, _ := getfullName()
	fmt.Println(firstName)
	// fmt.Println(middleName)
	// fmt.Println(lastName)
}
