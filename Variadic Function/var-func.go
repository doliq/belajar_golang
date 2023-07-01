package main

import (
	"fmt"
)

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(10, 20, 30, 40, 50)
	fmt.Println(total)

	numbers := []int{50, 40, 30, 20, 10}
	total = sumAll(numbers...)
	fmt.Println(total)
}
