package main

import "fmt"

type Biodata struct {
	Name, Age, Gender string
}

func main() {
	biodata1 := Biodata{"Hanif", "20", "men"}
	biodata2 := &biodata1

	biodata2.Name = "Indra"
	fmt.Println(biodata1) //Ikut berubah
	fmt.Println(biodata2) //Berubah menjadi Indra

	//menggunakan asterisk (*) agar ketika variabel yang menunjuk ke alamat yang sama akan
	//ikut berubah ketika salah satu variabel penunjuk ke alamat tersebut di alihkan ke alamat lain

	*biodata2 = Biodata{"Eva", "19", "women"}
	fmt.Println(biodata1)
	fmt.Println(biodata2)
}
