package main

import "fmt"

func main() {
	book := "The color of Magic"
	fmt.Println(book)

	fmt.Println(len(book))

	fmt.Println("book[0] = %v (type %T)\n", book[0], book[0])
}
