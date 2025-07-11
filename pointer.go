package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	/*
		address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
		address2 := address1 //copy value

		address2.City = "Bandung"
		fmt.Println("Pass by copy")
		fmt.Println(address1) // Ini tidak akan berubah
		fmt.Println(address2) // Ini akan mengalami perbuhan "Subang --> Bandung"
	*/

	//pointer

	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1 //pointer
	/*
		var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
		var address2 *Address = &address1 //pointer
	*/

	address2.City = "Bandung"
	fmt.Println("Pass by references")
	fmt.Println(address1) // Ini tidak akan berubah
	fmt.Println(address2) // Ini akan mengalami perbuhan "Subang --> Bandung"

}
