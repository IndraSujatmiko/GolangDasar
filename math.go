package main

import "fmt"

func main() {
	var (
		a = 10
		b = 20
		c = a + b
	)
	fmt.Println(c)

	var i = 10
	fmt.Println(i)

	i += 10
	fmt.Println(i)

	i += 5
	fmt.Println(i)

	var j = 8
	fmt.Println(j)
	j++
	fmt.Println(j)
	j++
	fmt.Println(j)

	j--
	fmt.Println(j)
	j--
	fmt.Println(j)

}
