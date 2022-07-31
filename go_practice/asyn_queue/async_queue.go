package asyn_queue

import (
	"context"
	"sync"
)

type CustomerWorker func(ctx context.Context, v ...interface{}) error //

// Config the common set
type Config struct {
	QueueSize     int // the size of queue
	CustomerCount int // 消费者的个数
}

// AsyncQueue the async queue struct
type AsyncQueue struct {
	ctx   context.Context  // ctx
	Name  string           // the name of the queue
	queue chan interface{} // queue
	wg    sync.WaitGroup   // control the goroutine process

	work       CustomerWorker
	cancelFunc context.CancelFunc // notice
	Config                        // the config of async queue
	startOnce  sync.Once          // stark once to do
	stopOnce   sync.Once          // stop once to do
}
