package main

import "fmt"

func sayGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	// initialize func as variable
	goodbye := sayGoodBye
	// the line below equal to sayGoodBye(param)
	// treat varible as function
	fmt.Println(goodbye("Eko"))
	fmt.Println(sayGoodBye("Eko"))
}
