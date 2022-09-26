package main

import "sync"

func main() {
	m := NewConcMap()

	m.Load(1, 2)
	println(m.Get(1))
}

type ConcMap struct {
	m map[int]int
	sync.RWMutex
}

func NewConcMap(size ...int) *ConcMap {
	if size != nil {
		return &ConcMap{m: make(map[int]int, size[0]), RWMutex: sync.RWMutex{}}
	} else {
		return &ConcMap{m: make(map[int]int), RWMutex: sync.RWMutex{}}
	}
}

func (c *ConcMap) Get(key int) (int, bool) {
	c.RLock()
	defer c.RUnlock()
	val, exists := c.m[key]
	return val, exists
}

func (c *ConcMap) Load(key int, val int) {
	c.Lock()
	defer c.Unlock()
	c.m[key] = val
}
