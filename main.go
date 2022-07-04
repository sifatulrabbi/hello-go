package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	// greet if there is a name otherwise skip greeting
	if len(args) > 1 && args[0] == "--name" {
		fmt.Println(greetings(args[1]))
	} else {
		fmt.Println("No name entered")
	}
}
