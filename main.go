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

// func findDuplicateId() {
// 	var (
// 		d          Ids
// 		duplicates int
// 		usedIds    []string
// 	)

// 	file := &lib.File{
// 		Path: FILE_PATH,
// 	}
// 	content, err := file.GetContent()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err := json.Unmarshal(content, &d); err != nil {
// 		log.Fatal(err)
// 	}

// 	for _, v := range d.Ids {
// 		for _, v2 := range usedIds {
// 			if v == v2 {
// 				duplicates += 1
// 			}
// 		}
// 		usedIds = append(usedIds, v)
// 	}
// 	fmt.Printf("Duplicates: %v\n", duplicates)
// }
