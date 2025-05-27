package main

import (
	"fmt"
	"time"
)

func PrintDurationWithTime(name string, iteration int, start time.Time) {
	duration := time.Since(start).Seconds()

	fmt.Printf("=== %-50s\n", name)
	fmt.Printf("[%d] iterations took seconds: [%f]\n\n", iteration, duration)
}

func PrintDuration(name string, iteration int, duration float64) {
	fmt.Printf("=== %-50s\n", name)
	fmt.Printf("[%d] iterations took seconds: [%f]\n\n", iteration, duration)
}
