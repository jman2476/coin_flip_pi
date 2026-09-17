package main

import "math/rand"

func flipCoin() bool {
	return rand.Float32() >= 0.5
}

func flipCoin64() bool {
	return rand.Uint64()&1 == 1
}

func flipCoin32() bool {
	return rand.Uint32()&1 == 1
}
