package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		var ch1 chan int
		ch1 <- 2
		fmt.Println(111)
		ch1 <- 1
		ch1 <- 3
		ch1 <- 4
		i := 0
		fmt.Println(i)
		i++
	}()
	time.Sleep(time.Second * 10)
	fmt.Println(1)
}
