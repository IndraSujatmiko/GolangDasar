package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

//Ini adalah Method

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var Indra Customer
	fmt.Println(Indra)

	Indra.Name = "Indra Sujatmiko"
	Indra.Address = "Lampung"
	Indra.Age = 20
	fmt.Println(Indra)
	fmt.Println(Indra.Address)

	//Cara Dua

	fatimah := Customer{
		Name:    "fatimah",
		Address: "Palembang",
		Age:     19,
	}
	fmt.Println(fatimah)

	//CARA TIGA

	gani := Customer{"Gani", "Jakarta", 23}
	fmt.Println(gani)

	//Memanggil method

	Indra.sayHello("Sarah")
}
