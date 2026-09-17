package main

import (
	"fmt"
	"sync"
)

func concurrentFlips(len int) string {
	results := make([]float64, len)
	var wg sync.WaitGroup

	for i := range len {
		wg.Add(1)

		go func(idx int) {
			defer wg.Done()
			results[idx] = flipProcess()
		}(i)
	}

	wg.Wait()

	var sum float64
	for _, v := range results {
		sum += v
	}

	pi_approx := sum / float64(len) * 4
	return fmt.Sprintf("Approximation of pi on %v iterations: %v", len, pi_approx)
}
