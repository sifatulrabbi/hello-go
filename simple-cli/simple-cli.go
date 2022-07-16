package simplecli

import (
	"fmt"
	"os"
	"strings"
)

func SimpleCLI() {
	args := os.Args[1:]
	// Showing help or mini tutorial
	if args[0] == "-h" || args[0] == "--help" || args[0] == "help" || args[0] == "h" {
		fmt.Printf(help)
		return
	}
	// Input: `app --name <name>`
	// Output: Your name is <name>
	if args[0] == "print" {
		print(args)
		return
	}
}

func print(args []string) {
	values := strings.Join(args[1:], " ")
	fmt.Printf(values + "\n")
}
