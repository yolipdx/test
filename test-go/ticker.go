package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.NewTicker(2 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case t := <- tick.C:
				fmt.Println("tick: ", t)
			case ok := <- done:
				if ok {
					return
				}
			}
		}
	}()

	time.Sleep(10 * time.Second)
	done <- true
}