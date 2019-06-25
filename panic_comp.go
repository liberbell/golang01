package main

import "fmt"

func safeValue(vals []int, index int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("ERROR: %s\n", err)
		}
	}()

	return vals[index]
}

func main() {
	v := safeValue([]int{1, 2, 3}, 10)
	fmt.Println(v)
}
