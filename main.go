package main

import (
	"fmt"
	"time"
)

func main() {
	iterations := []int{100, 1000, 10000, 20000, 30000, 40000, 50000}
	var results []string
	const timeLayout = "15:04:05.00000"

	for _, i := range iterations {
		timeStart := time.Now()
		fmt.Printf(
			"Starting %v iteration run @ %v\n",
			i, timeStart.Format(timeLayout))
		val := concurrentFlips(i)
		results = append(results, val)
		timeEnd := time.Now()
		fmt.Printf(
			"Finished %v iteration run @ %v\nResult: %v\n",
			i, timeEnd.Format(timeLayout), val)
		fmt.Printf(
			"Duration: %v\n", timeEnd.Sub(timeStart),
		)
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
