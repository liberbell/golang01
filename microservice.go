package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Print("Impliment me.")
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello Cloud Native Go.")
}
