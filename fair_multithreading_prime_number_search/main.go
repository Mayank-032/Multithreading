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

// 5760413
// 5761455

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

func doWork(wg *sync.WaitGroup, threadNum int) {
	defer wg.Done()
	start := time.Now()
	
	for {
		x := atomic.AddInt32(&currentNum, 1)
		if x > int32(maxInt) {
			break
		}
		checkPrime(int(x))
	}

	fmt.Printf("thread %s completed in %s\n", strconv.Itoa(threadNum), time.Since(start))
}

func main() {
	start := time.Now()
	
	var wg sync.WaitGroup
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go doWork(&wg, i)
	}

	wg.Wait()
	fmt.Println("checking till", maxInt, "found", totalPrimeNumbers+1, "prime numbers. Took", time.Since(start))
}