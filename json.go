package main

var data = `
{
  "user" : "Scrooge McDuck",
  "type" : "deposit",
  "amount" : 1000000.3,
}
`

type Request struct {
	Login  string  `json:"user"`
	Type   string  `json:"type"`
	Amount float64 `json:amount`
}
