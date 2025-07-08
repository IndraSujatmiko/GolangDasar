package main

import "fmt"

func getFullName() (string, string) {
	return "Indra", "Sujatmiko"
}

func nameSchool() (string, string, int) {
	return "sma", "negeri", 15
}

func main() {
	//firstName, lastName := getFullName()
	//fmt.Println(firstName, lastName)

	firstName, _ := getFullName()
	fmt.Println(firstName)

	nama1, nama2 := getFullName()
	fmt.Println(nama1, nama2)

	text1, text2, num := nameSchool()
	fmt.Println(text1, text2, num)
}
