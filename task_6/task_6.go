package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	/*
		1. Time.Sleep (time.After)
		2. context.Done
		3. chan msg received (chan closed)
		4. sync.cond.Signal (.Broadcast)
	*/

	cond := sync.NewCond(&sync.Mutex{})

	wg := sync.WaitGroup{}

	ctx, _ := context.WithTimeout(context.Background(), time.Second*4)

	finishChn := make(chan struct{})
	go func() {
		time.Sleep(time.Second * 6)
		close(finishChn)
		cond.Signal()
	}()

	GoSleep(time.Duration(2), &wg)
	GoCtxDone(ctx, &wg)
	GoChnMsg(finishChn, &wg)
	GoCondSignal(cond, &wg)

	wg.Wait()
}

func GoSleep(duration time.Duration, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * duration)
		fmt.Println("I finished!")
	}()
}

func GoCtxDone(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("I finished!")
				return
			default:
			}
		}
	}()
}

func GoChnMsg(chn <-chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-chn:
				fmt.Println("I finished!")
				return
			default:
			}
		}
	}()
}

func GoCondSignal(cond *sync.Cond, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		cond.L.Lock()
		defer cond.L.Unlock()
		cond.Wait()
		fmt.Println("I finished!")
	}()
}
