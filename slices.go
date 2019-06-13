package main

import "fmt"

func main() {
	loons := []string{"bugs", "duffy", "taz"}
	fmt.Printf("loons = %v (Type %T)\n", loons, loons)

	fmt.Println(len(loons))
	fmt.Println("-------------")
}
