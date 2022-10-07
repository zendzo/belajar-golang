package main

import "fmt"

func main(){
	var names [3]string
	names[0] = "Eko"
	names[1] = "Kurniawan"
	names[2] = "Khennedy"

	var values = [3]int {
		6,
		7,
		8,
	}
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(values)

	var arrayOut [7]int
	fmt.Println(len(arrayOut))
}
