package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("Yea")
	}

	if num := 10; num%2 == 0 {
		fmt.Println("Yes divisible")
	} else {
		fmt.Println("Not divisible")
	}
}
