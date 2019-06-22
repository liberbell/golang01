package main

import "fmt"

func NewTrade(Symbol string, Volume int, Price float64, Buy bool) (*Trade, error) {
	if Symbol == "" {
		return nil, fmt.Errorf("Symbol can`t be Empty")
	}

	if Volume <= 0 {
		return nil, fmt.Errorf("Volume must be >= 0 (was %d)", Volume)
	}

	if Price <= 0.0 {
		return nil, fmt.Errorf("Price must be >= 0(was %f)", Price)
	}

}
