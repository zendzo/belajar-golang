package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blacklist " + name)
	} else {
		fmt.Println("Welcome " + name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}

	blockRoot := func(name string) bool {
		return name == "root"
	}
	registerUser("admin", blacklist)
	registerUser("root", blockRoot)
	registerUser("Eko Kurniawan", blacklist)
}
