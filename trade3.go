package main

import "fmt"

func NewTrade(symbol string, volume int, price float64, buy bool) (*Trade, error) {
	if symbol == "" {
		return nil, fmt.Errorf("Symbol can`t be Empty")
	}

	if volume <= 0 {
		return nil, fmt.Errorf("Volume must be >= 0 (was %d)", volume)
	}

	if price <= 0.0 {
		return nil, fmt.Errorf("Price must be >= 0(was %f)", price)
	}

	trade := &Trade{
		Symbol: symbol,
		Volume: volume,
		Price:  price,
		Buy:    buy,
	}
	return trade, nil

}
