package main

import "sync"

func concurrentFlips(i int) float64 {
	results := make([]int, i)
	var wg sync.WaitGroup

}
