package backoff

import (
	"math/rand"
	"time"
)

type FibonacciBackOff struct {
	*ExponentialBackOff
	IntervalUnit time.Duration

	position uint
}

// Default values for FibonacciBackOff.
const (
	DefaultInitialPosition = 0
)

// NewFibonacciBackOff creates an instance of FibonacciBackOff using default values.
func NewFibonacciBackOff() *FibonacciBackOff {
	exp := NewExponentialBackOff()
	exp.InitialInterval = DefaultInitialInterval
	exp.RandomizationFactor = 0.0
	exp.Multiplier = 1.0
	exp.MaxInterval = DefaultMaxInterval
	exp.MaxElapsedTime = Infinity
	exp.Reset()

	fib := &FibonacciBackOff{
		ExponentialBackOff: exp,
		IntervalUnit:       time.Second,
		position:           DefaultInitialPosition,
	}
	fib.Reset()
	return fib
}

// Reset the interval back to the initial retry interval and restarts the timer.
// Reset must be called before using b.
func (b *FibonacciBackOff) Reset() {
	b.currentInterval = b.InitialInterval
	b.startTime = b.Clock.Now()

	b.position = DefaultInitialPosition
}

// NextBackOff calculates the next backoff interval using the formula:
// 	Randomized interval = RetryInterval * (1 ± RandomizationFactor)
func (b *FibonacciBackOff) NextBackOff() time.Duration {
	// Make sure we have not gone over the maximum elapsed time.
	if b.MaxElapsedTime != Infinity && b.GetElapsedTime() > b.MaxElapsedTime {
		return Stop
	}
	defer b.IncrementCurrentInterval()
	return getRandomValueFromInterval(b.RandomizationFactor, rand.Float64(), b.currentInterval)
}

// Increments the current interval with the fibonacciNumber * IntervalUnit.
func (b *FibonacciBackOff) IncrementCurrentInterval() {
	retryInterval := b.IntervalUnit * time.Duration(fibonacciNumber(b.position))
	b.position++

	// Check for overflow, if overflow is detected set the current interval to the max interval.
	if retryInterval >= b.MaxInterval {
		b.currentInterval = b.MaxInterval
	} else {
		b.currentInterval = retryInterval
	}
}

// fibonacciNumber calculates the Fibonacci sequence number for the given
// sequence position.
//
// 0, 1, 1, 2, 3, 5, 8 ...
func fibonacciNumber(n uint) uint {
	if 0 == n {
		return 0
	} else if 1 == n {
		return 1
	} else {
		return fibonacciNumber(n-1) + fibonacciNumber(n-2)
	}
}
