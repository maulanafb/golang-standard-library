package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(result)
	}

	// string ke integer
	resultInt, err := strconv.Atoi("1000")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(resultInt)
	}

	resultParse := strconv.FormatInt(999, 2)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(resultParse)
	}

	var string string = strconv.Itoa(999)
	fmt.Println(string)
}
