package main

type Point struct {
	X int
	Y int
}

func (p Point) Move(dx int, by int) {
	p.X += dx
	p.Y += dy
}

func main() {
	p := Point{1, 2}
}
