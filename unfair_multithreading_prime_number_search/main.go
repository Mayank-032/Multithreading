package main

import (
	"fmt"
	"math"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var maxInt = 100000000
var threads = 10
var totalPrimeNumbers int32 = 0
var currentNum int32 = 2

func checkPrime(x int) {
	if x&1 == 0 {
		return
	}

	for i := 3; i <= int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			return
		}
	}

	atomic.AddInt32(&totalPrimeNumbers, 1)
}

func createBatch(wg *sync.WaitGroup, threadNum, nStart, nEnd int) {
	defer wg.Done()
	start := time.Now()

	for i := nStart; i < nEnd; i++ {
		checkPrime(i)
	}

	fmt.Printf("thread %s from [%d - %d] completed in %s\n", strconv.Itoa(threadNum), nStart, nEnd, time.Since(start))
}

func main() {
	start := time.Now()
	nStart := 3
	batchSize := int(float64(maxInt) / float64(threads))

	var wg sync.WaitGroup
	for i := 0; i < threads-1; i++ {
		wg.Add(1)
		go createBatch(&wg, i, nStart, nStart+batchSize)
		nStart += batchSize
	}

	wg.Add(1)
	go createBatch(&wg, threads-1, nStart, maxInt)

	wg.Wait()
	fmt.Println("checking till", maxInt, "found", totalPrimeNumbers+1, "prime numbers. Took", time.Since(start))
}
