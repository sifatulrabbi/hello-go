package main

import (
	"github.com/sifatulrabbi/hello-go/pkg/lib"
)

const FILE_PATH string = "tmp/ids.json"

type Ids struct {
	Ids []string `json:"ids"`
}

func main() {
	lib.ProcessRibbons([]int{5, 2, 7, 4, 9}, 5)
	lib.ProcessRibbons([]int{1, 2, 5, 4, 9}, 4)
}
