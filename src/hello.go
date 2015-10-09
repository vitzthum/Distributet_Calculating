package main

import (
	"fmt"
	"time"
	"runtime"
	"sync"
)

const Threads = 10
var wg sync.WaitGroup
const Calculations = 1e10

type Vertex struct {
	start, end int
}


func guess(x Vertex) {
	defer wg.Done()

	for i := x.start; i <= x.end; i++ {
	}

	return
}

func main() {
	runtime.GOMAXPROCS(Threads)
	wg.Add(Threads)

	fmt.Println("")
	fmt.Print("type a value (0 - 1000): ")
	fmt.Scanln()

	var ranges [Threads]Vertex
	for i := 0; i < Threads; i++ {
		ranges[i].start = Calculations/Threads * i
		ranges[i].end = Calculations/Threads * (i+1) - 1
	}

	before := time.Now()

	for i := 0; i < Threads; i++ {
		go guess(ranges[i])
	}

	wg.Wait()
	duration := time.Since(before)

	fmt.Println("\n----------------------\n")
	fmt.Printf("Threads     : %v\n", Threads)
	fmt.Printf("Calculations: %v Billions\n", Calculations/1e9)
	fmt.Printf("done in     : %v\n", duration)
	fmt.Println("\n----------------------\n")

	// calculate average duration
//	averageDuration := calculateAverage(durations)

	// show average duration
//	showAverage(averageDuration, calculations)
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