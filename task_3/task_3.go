package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}

	//wg := sync.WaitGroup{}
	sumChn := make(chan int, len(arr))

	for _, val := range arr {
		//wg.Add(1)
		go func(val int) {
			//defer wg.Done()
			sumChn <- val * val
		}(val)
	}

	sum := 0
	// Вариант с wg.Wait
	//wg.Wait()
	//close(sumChn)
	//for val := range sumChn {
	//	sum += val
	//}

	// Вариант без wg.Wait
	for i := 0; i < len(arr); i++ {
		sum += <-sumChn
	}

	fmt.Println(sum)
}
