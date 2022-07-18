package simplecli

var help = `
Welcome to simpleCLI mini tutorial!

Usage: simpleCLI [COMMAND] [OPTIONS] [VALUE]...

COMMAND: print
This will print the next value/values your have entered.

COMMAND: minify
This will minify HTML, CSS, JavaScript files.
EXAMPLE: 
	simpleCLI minify -m html -f ./index.html -o ./minified/index.html
	simpleCLI minify -m js -f ./index.js
OPTIONS:
	-m, --method	Specify the method to use. Possible values html, css, js
	-f, --file	Path to he target file
	-o, --output	(optional) Output file path. If not specified it will flush the output in the same directory.

`
