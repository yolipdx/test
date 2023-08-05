package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func work(id int, ch chan bool) {
	fmt.Printf("work %d is running\n", id)

	for {
		select {
			case  _, ok := <- ch:
				if !ok {
					fmt.Printf("work %d detected chan closed, return\n", id)
					return
				}
			default:
				fmt.Printf("%d tick\n", id)
				time.Sleep(time.Second * time.Duration(rand.Int()%5))
		}
		
	}
	
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan bool)
	
	for i := 0; i< 10; i++ {
		wg.Add(1)
		id := i
		go func() {
			defer wg.Done()
			work(id, ch)
		}()
	}

	time.Sleep(10 * time.Second)
	close(ch)
	wg.Wait()
}