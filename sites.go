package main

import "fmt"

func returnType(url string) {
	resp, err := http.get(url)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
}
