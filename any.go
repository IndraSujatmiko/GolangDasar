package main

import "fmt"

//INTERFACE KOSONG (ANY)

func Hmm() any {
	return 1
}

func jago() any {
	return true
}

func main() {
	var kosong any = Hmm()
	fmt.Println(kosong)

	k := jago()
	fmt.Println(k)

}
