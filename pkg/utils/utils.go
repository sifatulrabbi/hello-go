package utils

import (
	"math/rand"
	"time"
)

func GetRandomNumber(min int32, max int32) (r int32) {
	rand.Seed(time.Now().UnixNano())
	r = rand.Int31n(max-min+1) + min
	return
}
