package main

import (
	"fmt"
//	"time"
	"runtime"
	"sync"
	"golang.org/x/crypto/bcrypt"
)

const Threads = 1
const Calculations = 1e6
var wg sync.WaitGroup
var found = false

type Vertex struct {
	start, end int
}


func guess(hash []byte) {
	defer wg.Done()

	start := []byte("A")
	end := []byte("z")

	var x []byte = make([]byte, 10)

	for i := 0; i <= 2 && !found; i++ {
		for j := start[0]; j <= end[0] && !found; j++ {
			x[i] = j
			if(bcrypt.CompareHashAndPassword(hash, x) == nil) {
				found = true
			}
		}
	}

	if found {
		fmt.Println("found :)")
	} else {
		fmt.Println("not found")
	}

	return
}

func main() {
	runtime.GOMAXPROCS(Threads)
	wg.Add(Threads)

	fmt.Println("")
	fmt.Print("type in your password (only 1 letter): ")

	var value string
	_, err := fmt.Scanln(&value)

	if(err != nil) {
		fmt.Println(err)
		return
	}

	password := []byte(value)
	hash, err := bcrypt.GenerateFromPassword(password, 4)

	if err != nil {
		fmt.Println(err)
		return
	}

/*
	var ranges [Threads]Vertex
	for i := 0; i < Threads; i++ {
		ranges[i].start = Calculations/Threads * i
		ranges[i].end = Calculations/Threads * (i+1) - 1
	}
	before := time.Now()
*/

	for i := 0; i < Threads; i++ {
		guess(hash)
	}

	wg.Wait()
/*
	duration := time.Since(before)

	fmt.Println("\n----------------------\n")
	fmt.Printf("Threads     : %v\n", Threads)
	fmt.Printf("Calculations: %v Millions\n", Calculations/1e6)
	fmt.Printf("done in     : %v\n", duration)
	fmt.Println("\n----------------------\n")

/*
	// calculate average duration
	averageDuration := calculateAverage(durations)

	// show average duration
	showAverage(averageDuration, calculations)
*/
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