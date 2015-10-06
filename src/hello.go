package main

import (
	"fmt"
	"time"
	"runtime"
	"sync"
)

const Threads = 10
var wg sync.WaitGroup

func calculate() () {
	defer wg.Done()

	for i := 1; i < 1e6/Threads; i++ {
		for j := 1; j < 1e5; j++ {
			var _ = i+j
		}
	}
	return
}
/*
func calculateAverage(durations [Threads]time.Duration) (averageDuration int) {
	for i := 0; i < Threads; i++ {
		averageDuration += int(durations[i])
	}
	averageDuration /= Threads

	time.ParseDuration(string(averageDuration))

	return
}

func showAverage(averageDuration int, calculations int) {
	fmt.Println("\n----------------------\n")
	if averageDuration/1e6 >= 1000 {
		fmt.Printf("average Time: %3.2f s\n", float64(averageDuration)/1e9)
	} else {
		fmt.Printf("average Time: %v ms\n", averageDuration/1e6)
	}
	fmt.Println("\n----------------------\n")
}
*/
func main() {
	runtime.GOMAXPROCS(Threads)
	wg.Add(Threads)

	before := time.Now()

	for i := 0; i < Threads; i++ {
		go calculate()
	}

	wg.Wait()
	fmt.Println("\n----------------------\n")
	fmt.Printf("done in: %v\n", time.Since(before))
	fmt.Println("\n----------------------\n")

	// calculate average duration
//	averageDuration := calculateAverage(durations)

	// show average duration
//	showAverage(averageDuration, calculations)
}