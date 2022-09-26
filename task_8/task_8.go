package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Print("Какое число изменяем: ")
	n := readInt64(in)

	fmt.Print("Позиция бита: ")
	pos := readInt64(in)

	val := strconv.FormatInt(n, 2)
	fmt.Println(val)
	valArr := []byte(val)

	changeBit(int(pos), 1, &valArr)

	val = string(valArr)
	fmt.Println(val)

}

func changeBit(bitPosition int, val byte, arr *[]byte) {
	if val == 0 {
		val = 48
	} else {
		val = 49
	}
	temp := *arr
	temp[bitPosition-1] = val
	arr = &temp
}

func readInt64(in *bufio.Reader) int64 {
	nStr, _ := in.ReadString('\n')
	nStr = strings.ReplaceAll(nStr, "\r", "")
	nStr = strings.ReplaceAll(nStr, "\n", "")
	n, _ := strconv.Atoi(nStr)
	return int64(n)
}

func readStr(in *bufio.Reader) string {
	nStr, _ := in.ReadString('\n')
	nStr = strings.ReplaceAll(nStr, "\r", "")
	nStr = strings.ReplaceAll(nStr, "\n", "")
	return nStr
}
