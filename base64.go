package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	value := "Muhammad Maulana Givari"
	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(decoded))
	}
}
