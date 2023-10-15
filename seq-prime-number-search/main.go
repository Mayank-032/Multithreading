package main

import (
	"fmt"
	"math"
	"time"
)

var maxInt = 100000000
var totalPrimeNumbers int32 = 0

func checkPrime(x int) {
	if x&1 == 0 {
		return
	}

	for i := 3; i <= int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			return
		}
	}

	totalPrimeNumbers++
}

func main() {
	start := time.Now()

	for i := 3; i < maxInt; i++ {
		checkPrime(i)
	}

	fmt.Println("checking till", maxInt, "found", totalPrimeNumbers+1, "prime numbers. Took", time.Since(start))
}