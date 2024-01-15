package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("ValidationError")
	NotFoundError   = errors.New("NotFoundError")
)

func main() {

	err := getById("")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("ValidationError")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("not found error")
		} else {
			fmt.Println("unknown error")
		}
	}
	// errors
}

func getById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "maulana" {
		return NotFoundError
	}
	// sukses
	return nil
}
