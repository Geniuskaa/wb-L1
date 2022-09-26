package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "snow dog sun"

	arrStr := strings.Split(str, " ")

	for i, j := 0, len(arrStr)-1; i < j; i, j = i+1, j-1 {
		arrStr[i], arrStr[j] = arrStr[j], arrStr[i]
	}

	str = strings.Join(arrStr, " ")
	fmt.Println(str)
}
