package main

import "fmt"

func main() {
	i := 12

	for i <= 20 {
		fmt.Println(i)
		i++
	}

	for j := 1; j <= 9; j++ {
		fmt.Println("Love")
	}

	for {
		fmt.Println("Yea")
		break
	}
}
