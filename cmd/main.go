package main

import (
	"github.com/sifatulrabbi/hello-go/pkg/lib"
)

func main() {
	f := lib.File{Path: "temp/test4.json"}
	// lib.CopyJSON("temp/test.json", "temp/test3.json")
	data := "\"name\": \"Sifatul\""
	lib.WriteJSON(f.Path, data)
}
