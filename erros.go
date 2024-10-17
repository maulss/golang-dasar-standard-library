package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "maulana" {
		return NotFoundError
	}

	return nil
}

func main() {
	//fmt.Println(getById("kocak"))

	err := GetById("budi")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Validation error", ValidationError)
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Not d=found error", NotFoundError)
		} else {
			fmt.Println("unknown error")
		}
	}
}
