package main

func NewSquare(x int, y int, length int) (*Square, error) {
  if length <= 0 {
    return nil, fmt.Errorf("Length must be > 0")
  }

  s := &Square {
    Center := Point{x, y},
    Length := length,
  }
  return s, nil
  
}
