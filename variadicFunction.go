package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(10, 10, 10, 10, 10, 10)
	fmt.Println(total)
	fmt.Println(sumAll(10, 10, 10))

	numbers := []int{20, 20, 20}
	fmt.Println(sumAll(numbers...))
}
