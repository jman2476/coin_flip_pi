package main

import (
	"fmt"
	"testing"
)

func TestMultiFlip(t *testing.T) {
	for range 5 {
		heads, total := multiFlipTest()
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

func TestHandleMultiFlip(t *testing.T) {
	for range 10 {
		head, tail, r := handleMultiFlip()

		if head <= tail {
			t.Errorf(
				"Heads is not greater than tails! \nH: %v, T: %v\n",
				head, tail,
			)
		}

		fmt.Printf("Ratio: %v, Ratio * 4: %v\n", r, r*4)
	}
}
