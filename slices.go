package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"John", "Paul", "George", "Ringo"}
	values := []int{1, 2, 3, 4, 5, 7, 10, 78}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "maulana"))
	fmt.Println(slices.Contains(values, 0))
	fmt.Println(slices.Index(names, "George"))
	fmt.Println(slices.Index(values, 0))

}
