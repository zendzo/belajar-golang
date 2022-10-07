package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perluangan", counter)
	// 	counter++
	// }

	// init statement
	// initiate value once, then check on every iteration and increase counter
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perluangan", counter)
	}

	var names = []string{"Muhammad", "Zaenal", "Mustoa"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// use _ to ignore index value
	for index, name := range names {
		fmt.Println("index ", index, "=", name)
	}

}
