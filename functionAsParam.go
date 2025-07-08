package main

import "fmt"

//CARA SATU
/*
func sayHelloWithFilter(name string, filter func(string) string) {
	fmt.Println("Hello", filter(name))
}
*/

//CARA DUA
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Indra", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}
