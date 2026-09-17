package main

import (
	"fmt"
	"math/rand"
)

func main() {
	iterations := []int{100, 1000, 10000, 100000, 1000000}
	var results []string

	for _, i := range iterations {
		val := concurrentFlips(i)
		results = append(results, val)
		fmt.Printf("Finished %v iteration run\nResult: %v\n", i, val)
	}

	for _, r := range results {
		fmt.Println(r)
	}
}

func getPi(i int) string {
	len := i
	var results []float64
	skips := 0

	for range len {
		ratio := flipProcess()

		results = append(results, ratio)
	}

	var sum float64
	for _, v := range results {
		sum += v
	}
	pi_approx := sum / float64(len) * 4

	return fmt.Sprintf("Approximation of pi on %v iterations, %v skips: %v", len, skips, pi_approx)
}

func flipProcess() float64 {
	var heads float64 = 0.0
	var tails float64 = 0.0

	for heads <= tails {
		if flipCoin() {
			heads++
		} else {
			tails++
		}
	}

	return heads / (tails + heads)
}

func flipCoin() bool {
	return rand.Float32() >= 0.5
}
