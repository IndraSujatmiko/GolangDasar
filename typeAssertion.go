package main

import "fmt"

func random() any {
	return 89
}

func main() {
	var result any = random()

	/*
		var resultString string = result.(string)
		fmt.Println(resultString)
		var resultInt int = result.(int)
		fmt.Println(resultInt)
	*/

	//type assertions --> reult.(string) <--Contoh

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}
}
