package main

import (
	"fmt"
	"testing"
)

func TestMultiFlip(t *testing.T) {
	for range 5 {
		heads, total := multiFlip()
		ratio := float64(heads) / float64(total)
		fmt.Printf(
			"Heads: %v, Total: %v, Ratio: %v\n",
			heads, total, ratio,
		)

		if ratio <= 0 {
			t.Errorf("Heads/Tails <= 0, Ratio = %v", ratio)
		}
	}
}
