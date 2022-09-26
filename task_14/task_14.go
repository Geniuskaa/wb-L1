package main

import (
	"fmt"
)

func main() {

	var a interface{}

	a = make(chan struct{})

	switch a.(type) {
	case int:
		fmt.Println("It`s int !")
	case string:
		fmt.Println("It`s string !")
	case bool:
		fmt.Println("It`s bool !")
	case chan struct{}:
		fmt.Println("It`s channel !")
	default:
		fmt.Println("Undefined type !")
	}
}
