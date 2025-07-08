package main

import "fmt"

func main() {

	/*
		name := "wahyu"

		if name == "Indra" {
			fmt.Println("Hai Indra")
		} else if name == "wahyu" {
			fmt.Println("Hai wahyu")
		} else {
			fmt.Println("bacot")
		}

		//SHORT STATEMENT

		if length := len(name); length > 5 {
			fmt.Println("Nama terlalu panjang")
		} else {
			fmt.Println("Nama sudah benar")
		}
	*/

	//konversi nilai siswa

	uas := 80
	uts := 90
	tugas := 100

	hasil := (uas * 50 / 100) + (uts * 30 / 100) + (tugas * 20 / 100)

	if hasil >= 80 {
		fmt.Println("Nilai akhir ", hasil, " (A)")
	} else if hasil > 65 && hasil < 80 {
		fmt.Println("Nilai akhir ", hasil, " (B)")
	} else {
		fmt.Println("Nilai akhir ", hasil, " (C)")
	}
}
