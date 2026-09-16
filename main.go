package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func main() {
	len := 100
	var results []float64
	skips := 0

	for range len {
		ratio, err := flipprocess()

		if err != nil {
			fmt.Println("Zero tails, skipped")
			len++
			skips++
			continue
		}
		results = append(results, ratio)
	}

	var sum float64
	for _, v := range results {
		sum += v
	}
	pi_approx := sum / float64(len) * 4

	fmt.Printf("Approximation of pi on %v iterations, %v skips: %v", len, skips, pi_approx)
}

func flipprocess() (float64, error) {
	var heads float64 = 0.0
	var tails float64 = 0.0

	for heads <= tails {
		if flipcoin() {
			heads++
		} else {
			tails++
		}
	}
	fmt.Printf("Heads: %v, Tails: %v\n", heads, tails)

	if tails == 0 {
		return 1.0, errors.New("Tails is zero")
	}

	return heads / tails, nil
}

func flipcoin() bool {
	return rand.Float32() >= 0.5
}
