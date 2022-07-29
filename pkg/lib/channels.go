package lib

import (
	"fmt"
	"time"
)

func Channels() {
	var (
		stream  = make(chan string)
		stream2 = make(chan string)
		done    = make(chan bool)
		done2   = make(chan bool)
	)

	go func() {
		// iCount := 100000
		for {
			stream <- "After 300ms"
			time.Sleep(time.Millisecond * 300)
		}
	}()
	go func() {
		for {
			stream2 <- "After 600ms"
			time.Sleep(time.Millisecond * 600)
		}
	}()
	go func() {
		for c := range stream {
			fmt.Println(c)
		}
		done <- true
	}()
	go func() {
		for c := range stream2 {
			fmt.Println(c)
		}
		done2 <- true
	}()
	go func() {
		for {
			time.Sleep(time.Second * 2)
			fmt.Println("This is an error string.")
		}
	}()

	<-done
	<-done2
	fmt.Println("Done!")
}
