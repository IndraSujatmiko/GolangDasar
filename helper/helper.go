package helper

var version = "2.0.0"
var Aplication = "Golang"

func sayGoodBye(name string) string {
	return "Good Bye " + name
}

func SayHello(name string) string { //nama function ditulis dengan awalan huruf besar, agar bisa di akses dari luar
	return "Hello " + name
}
