package example

import (
	"fmt"
	"time"
)

var ch1 = make(chan int, 2)
var ch2 = make(chan int, 2)
var done = make(chan struct{})

func SelectEx() {
	go receive()
	time.Sleep(time.Second * 3)
	ch1 <- 1
	ch2 <- 2
	<-done
}

func receive() {
	select {
	case a := <-ch1:
		fmt.Println("channel 1: ", a)

	case a := <-ch2:
		fmt.Println("channel 2: ", a)
	}

	done <- struct{}{}
}
