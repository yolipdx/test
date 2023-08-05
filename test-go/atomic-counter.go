package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func Inc(counter *uint64) {
	atomic.AddUint64(counter, 1)
}

func main() {
	var v uint64 = 0

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := 0; c < 100; c++ {
				Inc(&v)
			}
		}()
	}

	wg.Wait()
	fmt.Println("got v: ", v)
}
