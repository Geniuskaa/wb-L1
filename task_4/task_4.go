package main

import (
	"bufio"
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	fmt.Print("Сколько воркеров запускаем: ")
	in := bufio.NewReader(os.Stdin)

	countOfWorkers := readInt(in)
	wg := sync.WaitGroup{}

	chn := make(chan int)

	for i := 0; i < countOfWorkers; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("Worker number %d finished work\n", i)
					return
				case <-chn:
					fmt.Printf("Worker number %d begin work\n", i)
					continue
				}
			}

		}(i)
	}

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Main got Sigint signal!")
			wg.Wait()
			return
		default:
			chn <- rand.Intn(101)
			time.Sleep(time.Millisecond * 200)
		}
	}

}

func readInt(in *bufio.Reader) int {
	nStr, _ := in.ReadString('\n')
	nStr = strings.ReplaceAll(nStr, "\r", "")
	nStr = strings.ReplaceAll(nStr, "\n", "")
	n, _ := strconv.Atoi(nStr)
	return n
}
