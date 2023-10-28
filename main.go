package main

import (
	"fmt"

	"github.com/lloydChris/helloworld/stuff"
)

func main() {
	greeting := stuff.SayHello("Lorelei")
	fmt.Println(greeting)
}
