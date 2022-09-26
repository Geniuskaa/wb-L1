package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}

	wg := sync.WaitGroup{}

	for _, val := range arr {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			fmt.Printf("Square of %d is %d\n", val, val*val)
		}(val)
	}

	wg.Wait()
}
