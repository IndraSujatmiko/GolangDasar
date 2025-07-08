package main

import "fmt"

func main() {
	var name = "Indra Sujatmiko"
	fmt.Println(name)

	name = "Indra Jago"
	fmt.Println(name)

	lahir := "Lampung"   //untuk inisialisasi awal variabel tanpa harus menulis 'var' dapat diganti dengan ':='
	fmt.Println(lahir)

	lahir = "Bandar Lampung"
	fmt.Println(lahir)

	var (
		firstName = "Indra"
		lastName  = "Sujatmiko"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
