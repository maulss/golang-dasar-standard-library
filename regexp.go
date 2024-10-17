package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("edoo"))
	fmt.Println(regex.MatchString("eKoo"))

	fmt.Println(regex.FindAllString("eko edo ego, eto, efo", 10))

}
