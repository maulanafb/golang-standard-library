package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Eko")
	data.PushBack("Kurniawan")
	data.PushBack("Khannedy")

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
