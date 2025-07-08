package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16) // Pada kasus yang ini, nilai tidak muat karena terlalu besar sehingga akan kembali ke bawah yaitu -32768
	// jika nilai yang dimasukan 32769 artinya di melebihi dua angka yang seharusnya bisa diterma oleh tipe data int16 maka nilai yang dihasilkan adalah -32767

	var name = "Indra Sujatmiko"
	var I uint8 = name[0]
	var Istring = string(I)

	fmt.Println(name)
	fmt.Println(I)
	fmt.Println(Istring)
}
