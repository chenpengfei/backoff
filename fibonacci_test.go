package backoff

import (
	"testing"
	"time"
)

func TestFibonacciBackOff(t *testing.T) {
	var (
		testInitialInterval     = 500 * time.Millisecond
		testRandomizationFactor = 0.1
		testMultiplier          = 2.0
		testMaxInterval         = 8 * time.Second
		testMaxElapsedTime      = 15 * time.Minute
	)

	fib := NewFibonacciBackOff()
	fib.InitialInterval = testInitialInterval
	fib.RandomizationFactor = testRandomizationFactor
	fib.Multiplier = testMultiplier
	fib.MaxInterval = testMaxInterval
	fib.MaxElapsedTime = testMaxElapsedTime
	fib.IntervalUnit = time.Second
	fib.Reset()

	var expectedResults = []time.Duration{500, 0, 1000, 1000, 2000, 3000, 5000, 8000, 8000}
	for i, d := range expectedResults {
		expectedResults[i] = d * time.Millisecond
	}

	for _, expected := range expectedResults {
		assertEquals(t, expected, fib.currentInterval)
		// Assert that the next backoff falls in the expected range.
		var minInterval = expected - time.Duration(testRandomizationFactor*float64(expected))
		var maxInterval = expected + time.Duration(testRandomizationFactor*float64(expected))
		var actualInterval = fib.NextBackOff()
		if !(minInterval <= actualInterval && actualInterval <= maxInterval) {
			t.Error("error")
		}
	}
}
