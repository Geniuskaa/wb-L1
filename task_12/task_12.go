package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make([]string, 0, len(arr))

	m := make(map[string]struct{})

	for _, val := range arr {
		m[val] = struct{}{}
	}

	for k, _ := range m {
		set = append(set, k)
	}

	fmt.Println(set)
}
