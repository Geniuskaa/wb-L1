package main

import "fmt"

var justString string

func someFunc() {
	v := "10000000000"
	//Если мы решим указать правый предел, который превышает длину всей строки - мы получим "slice bounds out of range"
	justString = v[:10]
}

func main() {
	someFunc()
	fmt.Println(justString)
}
