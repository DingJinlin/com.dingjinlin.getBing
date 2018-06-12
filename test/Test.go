package test

import (
	"time"
	"fmt"
)

func Test() {
	now := time.Now()
	fmt.Println("start:",now)
	var sum int64 = 0
	var i int64 = 0

  //for ; i < 1000000000000; i++ {
	for ; i < 10000000000; i++ {
		sum = (sum + i)
	}

	duration := time.Since(now)
	fmt.Println("duration:", duration)
	fmt.Println("sum =", sum)
}

func TestFib()  {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}