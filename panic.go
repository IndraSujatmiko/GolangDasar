package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	message := recover() //Penggunaan recover pada posisi yang benar
	fmt.Println("terjadi panic", message)
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Ups Error")
	}

	/*
		CARA YANG SALAH DALAM MENGGUNKAAN RECOVER

		message := recover()
		fmt.Println("terjadi panic", message)
	*/
}

func main() {
	runApp(true)
	fmt.Println("Indra Sujatmiko") //jika salah dalam penggunaan recover maka yang ini tidak akan muncul
}
