package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Print("Impliment me.")
	http.ListenAndServe(":8080", nil)
}
