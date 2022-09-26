package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 8, 19, 23, 56, 70, 599}

	fmt.Println(findElem(arr, 19, 0))

}

func findElem(arr []int, target int, offset int) int {
	if len(arr) == 0 {
		return -1
	}

	pivot := len(arr) / 2

	if arr[pivot] > target {
		return findElem(arr[:pivot], target, offset)
	} else if arr[pivot] < target {
		return findElem(arr[pivot+1:], target, offset+pivot+1)
	} else if arr[pivot] == target {
		return pivot + offset
	}

	return -1
}
