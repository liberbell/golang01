package main

import "fmt"

func main() {
	n := 42
	s := fmt.Sprintf("%d", n)

	fmt.Printf("s = %s (Type %T)\n", s, s)
}
