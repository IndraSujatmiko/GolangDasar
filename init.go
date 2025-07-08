package main

import (
	"GolangDasar/database"
	_ "GolangDasar/internal" //Gunakan (_) ketika hanya ingin memangil method init
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}
