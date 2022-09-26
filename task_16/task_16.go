package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 5, 83, 8, 1, 6, 10}
	arr = quicksort(arr)
	fmt.Println(arr)
}

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	// Опорный элемент
	pivot := len(a) / 2

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}
