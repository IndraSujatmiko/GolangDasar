package main

import "fmt"

type Backlist func(string) bool

func registerUser(name string, backlist Backlist) {
	if backlist(name) {
		fmt.Println("You are blocked ", name)
	} else {
		fmt.Println("Welcome ", name)
	}
}

func main() {
	backlist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("anjing", backlist)

	registerUser("indra", func(name string) bool {
		return name == "anjing"
	})
}

// Kesimpulan : jika nilai yang dimasukan sama dengan return value name yang berisi anjing
// maka output akan mengeluarkan kalimat "You are blocked"