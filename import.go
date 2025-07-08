package main

import (
	"GolangDasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Indra")
	fmt.Println(result)

	// fmt.Println(helper.version)   // tidak bisa diakses karena huruf depan function kecil
	// fmt.Println(helper.sayGoodBye) // tidak bisa diakses karena huruf depan function kecil
	fmt.Println(helper.Aplication)
}
