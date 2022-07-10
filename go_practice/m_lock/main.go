package main

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	//addWithLock()
	//addWithNotLock()
	//LockCond()
	//Atom()
	//TestOnce()
	TestSync()
}

func TestSync() {
	var p sync.Pool
	p.Put(1)
	p.Put("a")
	fmt.Println(p.Get())
	fmt.Println(p.Get())

}

func TestWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Second * 3)
			fmt.Println("done!")
		}()
	}

	wg.Wait()
	fmt.Println("over")
}

func TestOnce() {
	var o sync.Once
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Second * 1)
			o.Do(func() {
				fmt.Println("mclink")
			})
		}()
	}
	wg.Wait()
	fmt.Println("over")
}

func Atom() {
	var a int64 = 10
	fmt.Println(atomic.AddInt64(&a, 1))
}

func LockCond() {
	// mailbox 代表信箱。
	// 0代表信箱是空的，1代表信箱是满的。
	var mailbox uint8
	// lock 代表信箱上的锁。
	var lock sync.RWMutex
	// sendCond 代表专用于发信的条件变量。
	sendCond := sync.NewCond(&lock)
	// recvCond 代表专用于收信的条件变量。
	recvCond := sync.NewCond(lock.RLocker())

	// sign 用于传递演示完成的信号。
	sign := make(chan struct{}, 3)
	max := 5
	go func(max int) { // 用于发信。
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= max; i++ {
			time.Sleep(time.Millisecond * 500)
			lock.Lock()
			for mailbox == 1 { // 还有信，可能会碰到“假唤醒”的情况。而且，
				// 如果存在“有多个wait但只需唤醒一个”的情况，也需要用for语句。在for语句里，
				//唤醒后可以再次检查状态，如果状态符合就开始后续工作，如果不符合就再次wait。用if语句就办不到
				sendCond.Wait() // 等待
			}
			log.Printf("sender [%d]: the mailbox is empty.", i)
			mailbox = 1
			log.Printf("sender [%d]: the letter has been sent.", i)
			lock.Unlock()
			recvCond.Signal() // 通知拿信
		}
	}(max)
	go func(max int) { // 用于收信。
		defer func() {
			sign <- struct{}{}
		}()
		for j := 1; j <= max; j++ {
			time.Sleep(time.Millisecond * 500)
			lock.RLock()
			for mailbox == 0 { //没有信阻塞
				recvCond.Wait()
			}
			log.Printf("receiver [%d]: the mailbox is full.", j)
			mailbox = 0
			log.Printf("receiver [%d]: the letter has been received.", j)
			lock.RUnlock()
			sendCond.Signal() // 通知放信
		}
	}(max)
	// 用来阻塞主协程，只有两个协程都完成了，调用了 sign <- struct{}{} ，主协程才会结束
	<-sign
	<-sign
}

func addWithNotLock() {
	var count int
	var wg sync.WaitGroup
	wg.Add(10) //添加10个等待
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done() //执行完一个协程去掉一个等待
			//对count进行10万次加1
			for j := 0; j < 100000; j++ {
				count++ //此处为临界区域
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

func addWithLock() {
	var count = 0 //公共修改的变量
	var mu = sync.Mutex{}
	//使用WaitGroup等待10个goroutine完成
	var wg sync.WaitGroup
	wg.Add(10) //添加10个等待
	//循环执行10个协程
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done() //执行完一个协程去掉一个等待
			//对count进行10万次加1
			for j := 0; j < 100000; j++ {
				//开始执行到临界区域 加锁
				mu.Lock()
				count++ //此处为临界区域
				//临界区域执行结束 释放锁
				mu.Unlock()
			}
		}()
	}
	//主线程等待10个协程完成
	wg.Wait()
	fmt.Println(count)
}
