package main

func main() {
	ch1, ch2 := make(chan int), make(chan int)

	go func() {
		ch1 <- 42
	}()
}
