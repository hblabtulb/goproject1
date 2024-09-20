package example

import "sync"

var cond = sync.NewCond(&sync.Mutex{})
var count = 0

func CondExample() {

	for i := 1; i < 10; i++ {
		go waitSometime(cond, i)
	}

	cond.L.Lock()

	cond.Wait()
	println(count)
}

func waitSometime(c *sync.Cond, index int) {

	count++
	cond.Signal()
}

func signal(c *sync.Cond, index int) {

}
