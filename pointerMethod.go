package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	indra := Man{"Indra"}
	indra.Married()

	fmt.Println(indra.Name)
}
