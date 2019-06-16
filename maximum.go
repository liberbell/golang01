package main

import "fmt"

func main() {
	nums := []int{16, 8, 42, 4, 23, 15}

	max := nums[0]

	for i := range nums {
		if max < nums[i] {
			max := nums[1]
		}
	}

	fmt.Println(max)
}
