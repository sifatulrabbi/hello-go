package lib

import (
	"fmt"
	"time"
)

func CountEverySecond() {
	current := time.Now()
	seconds := 0
	for {
		t := time.Now()
		if t.Unix()-current.Unix() > 1 {
			current = time.Now()
			seconds += 1
			fmt.Println(seconds)
		}
	}
}
