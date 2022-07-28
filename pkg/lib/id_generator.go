package lib

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var letters = [][]string{
	{"a", "d", "g", "j", "m", "p", "s", "v"},
	{"b", "e", "h", "k", "n", "q", "t", "w"},
	{"c", "f", "i", "l", "o", "r", "u", "x"},
}

func GetID() (id string) {
	t := time.Now().Unix()
	id = strconv.FormatInt(t, 10)
	// Convert the id string into a array of strings
	var idArr [10]string
	for i, v := range strings.Split(id, "") {
		if i < 10 {
			idArr[i] = v
		}
	}
	fmt.Println(idArr)

	return
}

func MakeLetterPairs() {
	r := getRandomNumber(0, 2)
	fmt.Println(r)
}

func getRandomNumber(min int32, max int32) (r int32) {
	rand.Seed(time.Now().UnixNano())
	r = rand.Int31n(max-min+1) + min
	return
}
