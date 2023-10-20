package main

import (
	"fmt"
	"sync"
	"time"
)

type ConcurrentQueue struct {
	queue []int32
}
var mutex sync.Mutex

func (q *ConcurrentQueue) Enqueue(wg *sync.WaitGroup, val int32) {
	mutex.Lock()
	q.queue = append(q.queue, val)
	mutex.Unlock()

	wg.Done()
}

func (q *ConcurrentQueue) Dequeue() int32 {
	val := q.queue[0]
	q.queue = q.queue[1:]
	return val
}

func (q *ConcurrentQueue) Size() int32 {
	return int32(len(q.queue))
}

func main() {
	q := ConcurrentQueue {
		queue: make([]int32, 0),
	}

	var wg sync.WaitGroup
	startTime := time.Now()
	
	/*
		// Description: Below operation is done one by one, but scenario is there are times when user comes together and update the queue so we must have a thread safe queue
		// 9+ micro-seconds time is taken and queue correctness is also ensured
		for i := 0; i < 1000000; i++ {
			q.Enqueue(int32(i))
		}
	*/

	/*
		// Description: Below operation is done concurrently, but it is a thread-unsafe-queue
		// Issue with this approach is correctness of queue is not ensured, sometime 900k, 915k etc.
		for i := 0; i < 1000000; i++ {
			wg.Add(1)
			go q.Enqueue(&wg, int32(i))
		}
	*/

	// Description: Below operation is done concurrently, and it is a thread-safe-queue
	// Issue with this qpproach is only performance as it is first acquiring the lock and remaining threads are still waiting to update/
	// but more than performance, initally correctness is more important
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go q.Enqueue(&wg, int32(i))
	}

	wg.Wait()

	fmt.Printf("TimeTaken: %v and totalQueueSize: %v\n", time.Since(startTime), q.Size())
}