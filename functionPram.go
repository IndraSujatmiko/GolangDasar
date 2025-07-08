package main

import "fmt"

func helloName(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	helloName("Indra", "Sujatmiko")
}
