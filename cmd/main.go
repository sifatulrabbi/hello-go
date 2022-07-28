package main

import (
	"github.com/sifatulrabbi/hello-go/pkg/lib"
)

func main() {
	lib.CopyJSON("temp/test.json", "temp/test3.json")
}
