package main

import (
	"fmt"
	"math/rand"
	"time"
)

func flipCoin() bool {
	return rand.Float32() >= 0.5
}

func flipCoin64() bool {
	return rand.Uint64()&1 == 1
}

func flipCoin32() bool {
	return rand.Uint32()&1 == 1
}

func testRandTimeSingle() {
	floatStart := time.Now()
	fltNum := rand.Float32()
	floatEnd := time.Now()

	uSixFourStart := time.Now()
	sixFourNum := rand.Uint64()
	uSixFourEnd := time.Now()

	uThreeTwoStart := time.Now()
	threeTwoNum := rand.Uint32()
	uThreeTwoEnd := time.Now()

	fmt.Printf(
		"Float %v took %v to compute\nuint64 %v took %v to compute\nuint32 %v took %v to compute\n",
		fltNum,
		floatEnd.Sub(floatStart),
		sixFourNum,
		uSixFourEnd.Sub(uSixFourStart),
		threeTwoNum,
		uThreeTwoEnd.Sub(uThreeTwoStart),
	)
}

func testFlipTimeSingle() {
	floatStart := time.Now()
	fltNum := flipCoin()
	floatEnd := time.Now()

	uSixFourStart := time.Now()
	sixFourNum := flipCoin64()
	uSixFourEnd := time.Now()

	uThreeTwoStart := time.Now()
	threeTwoNum := flipCoin32()
	uThreeTwoEnd := time.Now()

	fmt.Printf(
		"Float %v took %v to compute\nuint64 %v took %v to compute\nuint32 %v took %v to compute\n",
		fltNum, floatEnd.Sub(floatStart),
		sixFourNum,
		uSixFourEnd.Sub(uSixFourStart),
		threeTwoNum,
		uThreeTwoEnd.Sub(uThreeTwoStart),
	)
}

func multiFlip() (heads int, i int) {
	flips := rand.Uint64()
	for i = 0; i < 64; i++ {
		if (flips>>i)&1 == 1 {
			heads++
		}

		if heads > i+1-heads {
			return
		}
	}

	return
}

func handleMultiFlip(total int) {

}
