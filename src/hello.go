package main

import (
	"fmt"
	"time"
)

const Times = 5

func main() {

	var durationPerRound [Times]time.Duration
//	var averageTime int

	for i := 0; i < Times; i++ {

		fmt.Println("\n")
		fmt.Println("----------------------")
		fmt.Printf("\ncalculating: %v\n\n", i+1)

		x := 0

		bevore := time.Now()
		for i := 1; i < 100000; i++ {
			for j := 1; j < 8000; j++ {
				y := i+j
				if y >= 0 {
					x += 1
				}
			}
		}
		
		duration := time.Since(bevore)
		
		fmt.Printf("about %v Mio. calculations done in:  %v\n", x/1e6, duration)
		fmt.Println("\n----------------------\n")

		durationPerRound[i] = duration

		fmt.Println(duration.Hours)
	}

	for i := 0; i < Times; i++ {
//		averageTime += time.Nanoseconds(durationPerRound[i])
	}
//	averageTime /= Times

}
