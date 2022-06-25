package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var (
	amount int32 = 100
)

func a() {
	for i := 0; i < 1000; i++ {
		atomic.AddInt32(&amount, 10)
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("a done")
}

func b() {
	for i := 0; i < 1000; i++ {
		atomic.AddInt32(&amount, -10)
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("b done")
}

func main() {
	go a()
	go b()
	time.Sleep(1500 * time.Millisecond)
	fmt.Println("final amount: ", amount)
}
