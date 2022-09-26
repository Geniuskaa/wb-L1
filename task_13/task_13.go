package main

import "fmt"

func main() {
	a := 12
	b := 99

	a, b = b, a

	fmt.Println(a)
	fmt.Println(b)
}
