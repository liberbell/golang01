package main

import "os"

func main() {
	// vals := []int{1, 2, 3}
	// v := vals[10]
	// fmt.Println(v)
	file, err := os.Open("no-such-file")
	if err != nil {
		panic(err)
	}
}
