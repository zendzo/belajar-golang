package main

import "fmt"

func main(){
	person := map[string]string{
		"name" : "Eko",
		"address" : "Subang",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	delete(person, "title")
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Eko Kurniawan"
	book["wrong"] = "ups"

	delete(book, "wrong")
	fmt.Println(book)
}
