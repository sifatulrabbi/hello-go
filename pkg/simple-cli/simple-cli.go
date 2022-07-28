package simplecli

import (
	"fmt"
	"os"
	"strings"

	"github.com/sifatulrabbi/hello-go/pkg/simple-cli/services"
)

const INVALID_ERROR = "Invalid format. Please use 'help' or 'h' for helps.\n"
const NOT_FOUND_ERROR = "Command not found. Please use 'help' or 'h' to see help."

func SimpleCLI() {
	// All the cmd arguments
	var args []string = os.Args[1:]

	// Check if the user entered any commands or not.
	if len(args) < 1 {
		fmt.Print("No commands found\n")
		return
	}
	// Showing help or mini tutorial
	if args[0] == "-h" || args[0] == "--help" || args[0] == "help" || args[0] == "h" {
		fmt.Print(help)
		return
	}
	// Input: `simpleCLI print Hello`
	// Output: Hello>
	if args[0] == "print" {
		printAnyThing(args)
		return
	}
	// Minify HTML, JSON, CSS and JavaScript files
	if args[0] == "minify" {
		handleMinify(args)
	}
}

func printAnyThing(args []string) {
	values := strings.Join(args[1:], " ")
	fmt.Printf(values + "\n")
}

func handleMinify(args []string) {
	// Example: minify -m html -f ./index.html -o ./output.html
	if l := len(args); l < 5 {
		fmt.Print(INVALID_ERROR)
		return
	}

	// All the required vars
	cmd1, cmd2, method, path := args[1], args[3], args[2], args[4]

	// Check for the syntax
	if cmd1 != "-m" && cmd1 != "--method" || cmd2 != "-f" && cmd2 != "--file" {
		fmt.Print(INVALID_ERROR)
		return
	}
	// Find the method and run the minifier
	switch method {
	case "html":
		fmt.Print("Minifying html...\n")
		var minified string
		if len(args) > 5 {
			if len(args) != 7 {
				fmt.Print(INVALID_ERROR)
			} else {
				oPath := args[6]
				services.MinifyHTML(path, oPath)
			}
		} else {
			services.MinifyHTML(path, "")
		}
		fmt.Print(minified + "\n")
	case "css":
		fmt.Print("Minifying css...\n")
	case "js":
		fmt.Print("Minifying css...\n")
	default:
		fmt.Print(NOT_FOUND_ERROR)
	}
}
