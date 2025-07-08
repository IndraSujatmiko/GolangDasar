package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Indra"
	middleName = "Sujat"
	lastName = "miko"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
