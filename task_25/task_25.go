package main

import (
	"fmt"
	"time"
)

func main() {
	customSleep(time.Second * 10)
	fmt.Println("I waked up!")
}

func customSleep(t time.Duration) {
	<-time.After(t)
}
