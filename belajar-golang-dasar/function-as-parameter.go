package main

import (
	"fmt"
	"strings"
)

// use filter as paramter in declare function type as declaration
type Filter func(string) string

func sayHelloWithFiltered(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello " + nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "*censored*"
	} else {
		return name
	}
}

func upperCaseFilter(name string) string {
	filter := spamFilter
	return strings.ToUpper(filter(name))
}

func main() {
	filter := spamFilter
	sayHelloWithFiltered("Eko", filter)
	sayHelloWithFiltered("Anjing", spamFilter)
	sayHelloWithFiltered("Budi", upperCaseFilter)
}
