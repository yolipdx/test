package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(1000 * time.Millisecond)

	for {
		fmt.Println("iteration")
		select {
		case <-tick:
			fmt.Println("tick")
		case <- boom:
			fmt.Println("boom")
			return
		// default:
		// 	fmt.Println("           .")
		// 	time.Sleep(50 * time.Microsecond)
		}
	}
}