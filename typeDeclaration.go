package main

import "fmt"

func main() {

	type NoKtp string  // NoKtp adalah alias (kata lain) dari tipe data string, jadi ketika akan menggunakan string dapat menuliskan 'NoKtp'

	var ktpIndra NoKtp = "89898989898" // contoh penerapan

	var contoh string = "7979797979"
	var contohKtp NoKtp = NoKtp(contoh)

	fmt.Println(ktpIndra)
	fmt.Println(contohKtp)
}
