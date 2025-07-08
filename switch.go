package main

import "fmt"

func main() {
	name := "Indra"

	switch name {
	case "Indra":
		fmt.Println("Hello Indra")
	case "Imelda":
		fmt.Println("Hello Imelda")
	default:
		fmt.Println("Boleh Kenalan..?")
	}

	//SHORT STATEMENT

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
