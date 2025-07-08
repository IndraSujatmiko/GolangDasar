package main

import "fmt"

func main() {

	//MEMBUAT SLICE DENGAN ARRAY YANG SUDAH ADA

	names := [...]string{"Eko", "Indra", "Jamilah", "Sukma", "Sefira", "Dimas"}

	slice := names[4:6] //Mengambil data dengan index 4 sampai sebelum 6
	fmt.Println(slice)

	slice2 := names[:3] //Mengambil data dengan index kurang dari 3
	fmt.Println(slice2)

	slice3 := names[3:] //Mengambil data dengan index lebih dari 3
	fmt.Println(slice3)

	slice4 := names[:] //Mengambil semua data
	fmt.Println(slice4)

	/*
		contoh slice :
		var slice []string = names[:]
		fmt.Println(slice)
	*/

	//APPEND SLICE (MENAMBAH DATA)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println("Ini adalah hari setelah di perbarui : ", days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Ups"
	fmt.Println("Ini adalah data hari pada Slice 2 : ", daySlice2) //Karena pada array days sudah tidak lagi dapat menampung data maka data yang kita tabahkan ini akan secara otomatis dimasukan pada array baru yang dibuat otomatis, namun array yang dimaksud tidak dapat kita ketahui
	fmt.Println("Ini adalah data hari dari sebuah array : ", days)
	fmt.Println("Ini adalh data hari dari Slice 1 : ", daySlice1)

	//MAKE SLICE

	newSlice := make([]string, 2, 5) //newSlice []string := make([]string, panjang, kapasitas)
	newSlice[0] = "Indra"
	newSlice[1] = "Sujat"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Miko")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "hehe"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	//COPY SLICE

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	/*
		iniArray := [...]int{1,2,3,4,5}
		iniSlice := []int{1,2,3,4,5}
	*/
}
