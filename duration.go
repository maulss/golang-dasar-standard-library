package main

import (
	"fmt"
	"time"
)

func main() {
	var duration1 time.Duration = 100 * time.Second
	var duration2 time.Duration = 10 * time.Millisecond
	var duration3 time.Duration = duration1 - duration2

	fmt.Println(duration1)
	fmt.Println(duration2)
	fmt.Println(duration3)

	fmt.Printf("%.2f\n", duration1.Seconds())
	fmt.Printf("Duration : %d\n", duration3)
}
