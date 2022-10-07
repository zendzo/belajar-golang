package main

import "fmt"

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

// factorialRecrusive
func factorialRecrusive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecrusive(value-1)
	}
}

func main() {
	result := factorialLoop(5)
	fmt.Println(result)
	fmt.Println(factorialRecrusive(5))
}
