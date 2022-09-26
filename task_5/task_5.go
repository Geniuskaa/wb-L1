package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	chn := make(chan int, 1)
	finishChn := make(chan struct{})

	fmt.Print("Сколько секунд работаем: ")
	in := bufio.NewReader(os.Stdin)
	n := readInt(in)

	//go func() {
	//	for {
	//		_, open := <-chn
	//		if !open {
	//			fmt.Println("Goroutine is finishing")
	//			return
	//		}
	//		fmt.Println("Прочитали")
	//	}
	//}()

	go func() {
		time.Sleep(time.Second * time.Duration(n))
		finishChn <- struct{}{}
	}()

	for {
		select {
		case <-finishChn:
			close(finishChn)
			close(chn)
			fmt.Println("Main finished work")
			return
		default:
			select {
			case chn <- 2:
				fmt.Println("Записали")
			case <-chn:
				fmt.Println("Прочитали")
			}
		}

		time.Sleep(time.Millisecond * 400)
	}

}

func readInt(in *bufio.Reader) int {
	nStr, _ := in.ReadString('\n')
	nStr = strings.ReplaceAll(nStr, "\r", "")
	nStr = strings.ReplaceAll(nStr, "\n", "")
	n, _ := strconv.Atoi(nStr)
	return n
}
