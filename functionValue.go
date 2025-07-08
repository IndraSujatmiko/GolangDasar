package main

import "fmt"

func getGoodbye(name string) string {
	return "GoodBye " + name
}

func main() {

	contoh := getGoodbye //Ini intinya
	fmt.Println(contoh("Indra"))
}
