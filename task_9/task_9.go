package main

import (
	"fmt"
)

func main() {
	startChn := make(chan int)
	finisgChn := make(chan int)

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	go func() {
		for _, val := range arr {
			startChn <- val
		}
		close(startChn)
	}()

	go func() {
		for val := range startChn {
			finisgChn <- val * 2
		}
		close(finisgChn)
	}()

	for val := range finisgChn {
		fmt.Println(val)
	}
}
