package main

import "fmt"

func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

func main() {
	result := getHello("Indra")
	fmt.Println(result)

	for x := 0; x < 10; x++ {
		fmt.Println(result)
	}

	fmt.Println(getHello("Hana"))
}
