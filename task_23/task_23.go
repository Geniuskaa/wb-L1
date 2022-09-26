package main

import (
	"errors"
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6}

	delElemOnPos := 2

	a, err := delElem(a, delElemOnPos)
	if err == nil {
		fmt.Println(a)
	}

}

func delElem(arr []int, pos int) ([]int, error) {
	if pos > len(arr) {
		return nil, errors.New("Your positon more than length of slice")
	}

	arr = append(arr[:pos-1], arr[pos:]...)
	return arr, nil
}
