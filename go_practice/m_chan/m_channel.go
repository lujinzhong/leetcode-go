package main

import (
	"log"
	"sync"
	"time"
)

var active = make(chan struct{}, 3)
var jobs = make(chan int, 10)

func main() {
	go func() {
		for i := 0; i < 8; i++ {
			jobs <- i + 1
		}
		close(jobs) // 让 for 循环退出
	}()
	time.Sleep(time.Second * 3)
	var wg sync.WaitGroup
	for j := range jobs {
		wg.Add(1)
		go func(j int) {
			active <- struct{}{} // 有三个正在消费的话，会阻塞
			log.Printf("handle job: %d\n", j)
			time.Sleep(2 * time.Second)
			<-active
			wg.Done()
		}(j)
	}
	wg.Wait()
}
