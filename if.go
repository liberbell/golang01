package main

import "fmt"

func main() {
	x := 10

	if x > 5 {
		fmt.Println("x is big")
	}
	if x > 100 {
		fmt.Println("x is very big")
	} else {
		fmt.Println("x is not that big")
	}

	if x > 5 && x < 15 {
		fmt.Println("x is just right")
	}

	if x < 20 || x > 30 {
		fmt.Println("x is out of range")
	}
}
