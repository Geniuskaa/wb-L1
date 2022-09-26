package main

import "fmt"

func main() {
	a := []int{1, 6, 3, 93, 982, 1, 47, 2, 7}
	b := []int{2, 7, 3, 63, 92, 33, 5, 72, 37}
	c := make([]int, 0, len(a)+len(b))

	m := make(map[int]int, len(a))

	for _, val := range a {
		m[val]++
	}

	for _, val := range b {
		_, exists := m[val]
		if exists {
			c = append(c, val)
		}
	}

	fmt.Println(c)
}
