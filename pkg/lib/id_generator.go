package lib

import (
	"strconv"
	"strings"
	"time"

	"github.com/sifatulrabbi/hello-go/pkg/utils"
)

var letters = [][]string{
	{"aQ", "Op", "go", "vJ", "Am", "pC", "sH", "Bv", "wp", "sd"},
	{"Ea", "ji", "Hl", "Sk", "nZ", "nM", "tS", "xw", "su", "qw"},
	{"Cx", "iD", "sI", "Ll", "oP", "Qr", "xu", "Xd", "Ap", "LL"},
}
var prevIds []string // Slice of all the used ids.

func GetID() (id string) {
	var duplicate bool
	id = buildID()
	// Validate the id before returning
	duplicate = validateID(id)
	if !duplicate {
		prevIds = append(prevIds, id)
	}
	// for !duplicate {
	// 	id = buildID()
	// 	duplicate = validateID(id)
	// }
	return
}

func buildID() (id string) {
	t := time.Now().Unix()
	id = strconv.FormatInt(t, 10)
	// Convert the id string into a array of strings
	var idArr []string
	for _, v := range strings.Split(id, "") {
		n, _ := strconv.ParseInt(v, 10, 64)
		lArr := letters[utils.GetRandomNumber(0, 2)]
		randStr := lArr[n]
		idArr = append(idArr, randStr)
	}
	id = strings.Join(idArr, "")
	return
}

// Check if the id is already used or not
func validateID(id string) bool {
	for _, p := range prevIds {
		if p == id {
			return true
		}
	}
	return false
}
