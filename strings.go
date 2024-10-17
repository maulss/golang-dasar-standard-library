package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Maulana Givari", "Mau"))
	fmt.Println(strings.Split("Maulana Givari", " "))
	fmt.Println(strings.ToLower("Maulana Givari"))
	fmt.Println(strings.ToUpper("Maulana Givari"))
	fmt.Println(strings.Trim("   Maulana Givari   ", " "))
	fmt.Println(strings.ReplaceAll("Maulana Muhammad Maulana Givari", "Maulana", "GG"))
}
