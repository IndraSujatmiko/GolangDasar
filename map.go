package main

import "fmt"

func main() {
	/*
		CARA LAIN MEMBUAT TIPE DATA MAP

		var person map[string]string = map[string]string{}
		person["name"] = "Indra"
		person["address"] = "Lampung"
	*/

	person := map[string]string{
		"name":    "Indra",
		"address": "Lampung",
	}

	//CARA MEMANGGIL DATA PADA TIPE DATA MAP

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	//MAP BARU

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Indra"
	book["hmm"] = "salah"

	fmt.Println(book)

	delete(book, "hmm")

	fmt.Println(book)
}
