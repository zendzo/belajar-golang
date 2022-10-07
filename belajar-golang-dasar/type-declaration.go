package main

import "fmt"

func main(){
	type noKTP string
	type marriedStatus bool
	var isMarried marriedStatus = true
	var idKTP noKTP = "777777777777"

	fmt.Println(isMarried)
	fmt.Println(idKTP)
}

