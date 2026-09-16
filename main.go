package main

import (
	"fmt"
	"math/rand"
)

func main() {
	iterations := []int{100, 1000, 10000, 100000}
	var results []string

	for _, i := range iterations {
		results = append(results, getpi(i))
		fmt.Printf("Finished %v iteration run\n", i)
	}

	for _, r := range results {
		fmt.Println(r)
	}
}

func getpi(i int) string {
	len := i
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

	return fmt.Sprintf("Approximation of pi on %v iterations, %v skips: %v", len, skips, pi_approx)
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
	// fmt.Printf("Heads: %v, Tails: %v\n", heads, tails)

	// if tails == 0 {
	// 	return 1.0, errors.New("Tails is zero")
	// }

	return heads / (tails + heads), nil
}

func flipcoin() bool {
	return rand.Float32() >= 0.5
}
