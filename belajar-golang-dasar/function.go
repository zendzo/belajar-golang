package main

import "fmt"

func sayHello(name string) {
	fmt.Println("Hello " + name)
}

func main() {
	var sliceNames = make([]string, 2, 10)
	sliceNames[0] = "Budi Rahardi"
	sliceNames[1] = "Rian Rahardi"

	for _, name := range sliceNames {
		sayHello(name)
	}

	fmt.Println(cap(sliceNames))
}
