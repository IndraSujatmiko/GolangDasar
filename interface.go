package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string { //implementasi dari struct --> HasName
	return person.Name
}

func main() {
	person := Person{Name: "Indra"}
	sayHello(person)
}
