package main

import "fmt"

func main() {
	tempArr := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	m := make(map[int][]float32)

	for _, val := range tempArr {
		c := int(val / 10)

		v, exists := m[c*10]
		if exists {
			v = append(v, val)
			m[c*10] = v
		} else {
			m[c*10] = []float32{val}
		}
	}

	fmt.Println(m)

}
