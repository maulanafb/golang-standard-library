package main

import (
	"fmt"
	"strings"
)

type Man struct {
	Name string
}

func main() {
	fmt.Println("Starting")
	Maulana := Man{"Maulana Fatih"}

	fmt.Println(Maulana.toLower())
	fmt.Println(ToUpper(&Maulana))
	fmt.Println(strings.Split(Maulana.Name, " "))
}

func (man *Man) toLower() string {
	return strings.ToLower(man.Name)
}

func ToUpper(man *Man) string {
	return strings.ToUpper(man.Name)
}
