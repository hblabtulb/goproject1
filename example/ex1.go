package example

import (
	"fmt"
	"log"
	"time"
)

var ch = make(chan int, 2)

func Ex1() {

	fmt.Printf("Hello and welcome to %s!\n", "Hblab")
	//go HelloWorld(1)
	//go HelloWorld(2)
	//go HelloWorld(3)

	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)

	go receiver()
	ch <- 1
	ch <- 2
	//for n, bo := range ch {
	//	fmt.Printf("from: %d, bo: %s", n, bo)
	//}
	close(ch)

	n, bo := <-ch
	fmt.Printf("1 - from: %d, bo: %t\n", n, bo)

	//n, bo = <-ch
	//fmt.Printf("2- from: %d, bo: %t\n", n, bo)

	//n, bo = <-ch
	//fmt.Println("=================")
	//fmt.Printf("3 - from: %d, bo: %t\n", n, bo)

}

func receiver() {
	n, bo := <-ch
	fmt.Printf("receiver: %d, bo: %t\n", n, bo)
}

func HelloWorld(id int) {
	time.Sleep(1000 * time.Millisecond)
	log.Println("Hello, World!: ", id)

	ch <- id
	if id == 3 {
		close(ch)
	}
}
