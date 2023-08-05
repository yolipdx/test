package main

import (
	"fmt"
	"sync"
	"time"
	// "time"
)

type counter struct {
	value int
}


func inc(c *counter) {
	c.value += 1
}

func main_1() {
	cnt := counter{0}

	for i:=0; i<1000; i++ {
		go inc(&cnt)
	}


	time.Sleep(time.Second * 2)
	fmt.Println("cnt: ", cnt.value)
}


func main() {
	cnt := counter{0}

	var lock sync.Mutex

	for i:=0; i<1000; i++ {
		go func(c *counter) {
			lock.Lock()

			c.value += 1

			lock.Unlock()

		}(&cnt)
	}


	time.Sleep(time.Second * 2)
	lock.Lock()
	fmt.Println("cnt: ", cnt.value)
	lock.Unlock()
}