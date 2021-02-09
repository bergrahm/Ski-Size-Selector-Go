package main

import (
	"flag"
	"fmt"
)

type person struct {
	name  string
	age   int
	style string
}

func initPerson() person {
	return person{}
}
func main() {
	flag.Parse()
	fmt.Println("Hello, world.")
}
