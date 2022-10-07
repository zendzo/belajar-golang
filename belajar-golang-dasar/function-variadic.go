package main

import "fmt"

func sumAll(nums ...int) int {
	total := 0

	for _, value := range nums {
		total += value
	}
	return total
}

func main() {
	total := sumAll(10, 20, 30, 40)
	fmt.Println(total)

	slice := []int{10, 20, 30, 40}
	total = sumAll(slice...)
	fmt.Println(total)
}
