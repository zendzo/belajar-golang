package main

import "fmt"

func main() {
	var name = "Eko"

	if name == "Eko" {
		fmt.Println("Eko")
	} else if name == "Budi" {
		fmt.Println("Budi")
	} else {
		fmt.Println("no one")
	}
	// short statement, check
	// only availabe in golang
	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama Benar")
	}
}
