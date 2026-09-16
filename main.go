package main

import (
	"fmt"
	"math/rand"
)

func main() {
	len := 100
	var results []float64

	for range len {
		results = append(results, flipprocess())
	}

	var sum float64
	for _, v := range results {
		sum += v
	}
	pi_approx := sum / float64(len) * 16

	fmt.Printf("Approximation of pi on %v iterations: %v", len, pi_approx)
}

func flipprocess() float64 {
	var heads float64 = 0.0
	var tails float64 = 0.0

	for heads <= tails {
		if flipcoin() {
			heads++
		} else {
			tails++
		}
	}

	return heads / tails
}

func flipcoin() bool {
	return rand.Float32() >= 0.5
}
