package main

import "fmt"

func sayHelloTo(firstName, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	name := "Eko"
	sayHelloTo(name, "Kurniawan")
	sayHelloTo("Budi", "Setiawan")
}
