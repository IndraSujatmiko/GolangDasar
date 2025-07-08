package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var nilaiAbsensi = 80

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = nilaiAbsensi > 80

	var lulus bool = lulusAbsensi && lulusNilaiAkhir

	fmt.Println(lulus)
}
