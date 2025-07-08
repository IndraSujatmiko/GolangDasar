package main

import "fmt"

func main() {
	var names [2]string
	names[0] = "Indra"
	names[1] = "Sujatmiko"

	fmt.Println(names[0])
	fmt.Println(names[1])

	var values = [...]int{
		90,
		89,
		79,
		69,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[3])
	fmt.Println(values[2])
	values[2] = 100 //Mengubah isi data array
	fmt.Println(values[2])
	fmt.Println(len(values))
}
