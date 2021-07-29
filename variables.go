package main

import "fmt"

func main() {
	var a = "shahrukh"
	fmt.Println(a)
	z := "cool"

	fmt.Println(z)
	var b, c int = 1, 2

	fmt.Println(b, c)

	var d, e = 3, 4

	fmt.Println(d, e)

	var f = true // type inferred automatically

	fmt.Println(f)

	var g int

	fmt.Println(g) // g initialized to nil value of zero for int
}
