package main

import "fmt"

func main() {
	var name = "M Zaenal Must"

	switch name {
	case "Eko":
		fmt.Println("Halo Eko")
	case "Budi":
		fmt.Println("Halo Budi")
	default:
		fmt.Println("Kenalan Donk")
	}

	// short statement switch
	// checking condition before expresion
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// switch without expresion
	// also used as subtitue simple if condition

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
