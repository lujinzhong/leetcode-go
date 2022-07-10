package task_queue

import (
	"fmt"
	"sync"
	"testing"
)

func TestQueue(t *testing.T) {
	res, err := divide(1, 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func divide(a, b int) (res int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("omg, panic ! err:%v", r)
			return
		}
	}()
	res = a / b
	return
}

func add(res, a int) int {
	var mu sync.Mutex
	mu.Lock()
	res += a
	mu.Unlock()
	return res
}
