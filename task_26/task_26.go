package main

import (
	"fmt"
	"strings"
)

func main() {
	arr := []string{"abcd", "abCdefAaf", "aabcd"}

	for _, val := range arr {
		val = strings.ToLower(val)
		m := make(map[rune]struct{})
		okay := true

		for _, vv := range val {
			if _, exists := m[vv]; exists {
				okay = false
				break
			} else {
				m[vv] = struct{}{}
			}
		}

		if okay {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
	}
}
