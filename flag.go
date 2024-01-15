package main

import (
	"flag"
	"fmt"
)

func main() {

	var username *string = flag.String("username", "root", "database username")
	var password *string = flag.String("password", "root", "database password")
	var host *string = flag.String("host", "localhost", "database host")
	var port *int = flag.Int("port", 0, "database port")

	flag.Parse()

	fmt.Println(*username)
	fmt.Println(*password)
	fmt.Println(*host)
	fmt.Println(*port)
}
