package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type counter struct {
	n atomic.Int32
}

func newCounter() *counter {
	return &counter{n: atomic.Int32{}}
}

func (c *counter) Inc() {
	c.n.Add(1)
}

func main() {
	c := newCounter()
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}

	wg.Wait()
	fmt.Println(c.n.Load())
}
