package main

import "fmt"

func main() {
	number := 1

	for number <= 10 {
		fmt.Println("Perulangan ke ", number)
		number++
	}
	fmt.Println("Selesai")

	//FOR DENGAN STATEMENT

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}
	fmt.Println("Ini dengan statement")

	// FOR 2 Manual

	names := []string{"Indra", "Karin", "Sukma"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// FOR RANGE

	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}

	for _, name := range names {
		fmt.Println(name)
	}

}
