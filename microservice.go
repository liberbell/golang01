package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Print("Impliment me.")
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

}
