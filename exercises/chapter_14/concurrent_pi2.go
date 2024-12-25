package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

const NCPU = 2

func term2(ch chan float64, start, end int) {
	result := 0.0
	for i := start; i < end; i++ {
		x := float64(i)
		result += 4 * (math.Pow(-1, x) / (2.0*x + 1.0))
	}
	ch <- result
}

func CalculatePi2(end int) float64 {
	ch := make(chan float64)
	for i := 0; i < NCPU; i++ {
		go term2(ch, i*end/NCPU, (i+1)*end/NCPU)
	}

	result := 0.0
	for i := 0; i < NCPU; i++ {
		result += <-ch
	}
	return result
}

func mainConcurrentPi2() {
	start := time.Now()
	runtime.GOMAXPROCS(2)
	fmt.Println(CalculatePi2(5000))
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}
