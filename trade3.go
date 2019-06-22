package main

import "fmt"

func NewTrade(Symbol string, Volume int, Price float64, Buy bool) (*Trade, error) {
	if Symbol == "" {
		return nil, fmt.Errorf("Symbol can`t be Empty")
	}
}
