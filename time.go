package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Println(now.Local())

	utc := time.Date(2021, time.April, 16, 1, 2, 36, 0, time.UTC)
	fmt.Println(utc.Local())

	formatter := "2006-01-02 15:04:05"

	value := "2020-10-10 10:10:10"

	result, err := time.Parse(formatter, value)

	if err != nil {
		fmt.Println("Error :", err.Error())
	} else {
		fmt.Println(result)
	}
}
