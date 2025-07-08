package main

import "fmt"

func factorialLoop(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecrusive(value int) int {
	if value == 1 {
		return value
	} else {
		return value * factorialRecrusive(value-1)
	}
}

func main() {
	fmt.Println("Ini menggunakan reqursive : ", factorialRecrusive(10))
	fmt.Println("Ini menggunkan for loop : ", factorialLoop(10))
}
