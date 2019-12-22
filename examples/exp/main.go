package main

import (
	"fmt"
	"github.com/chenpengfei/backoff"
	"time"
)

func main() {
	fib := backoff.NewExponentialBackOff()
	fib.MaxInterval = time.Hour
	fib.MaxElapsedTime = backoff.Infinity

	for {
		fmt.Printf("interval: %v\n", fib.NextBackOff())
		time.Sleep(500 * time.Millisecond)
	}
}
