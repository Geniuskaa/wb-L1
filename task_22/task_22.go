package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := newBigInt("20000000000000000000000000000000000000000000100")
	b := newBigInt("2000000000000000000000000000000000000000000500")

	mul := &big.Int{}
	mul.Mul(a, b)
	fmt.Println(mul.String())

	div := &big.Int{}
	div.Div(a, b)
	fmt.Println(div.String())

	add := &big.Int{}
	add.Add(a, b)
	fmt.Println(add.String())

	sub := &big.Int{}
	sub.Sub(a, b)
	fmt.Println(sub.String())
}

func newBigInt(bigNum string) *big.Int {
	a := &big.Int{}

	a.SetString(bigNum, 10)

	return a
}
