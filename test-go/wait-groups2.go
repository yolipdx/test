package main

import (
	"fmt"
	"sync"
)

func worker(jobId int, wg *sync.WaitGroup) {
	fmt.Printf("job %d running\n", jobId)
	
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	jobNum := 3
	
	for i := 0; i< jobNum; i++ {
		id := i
		wg.Add(1)
		

		go worker(id, &wg)
	}

	fmt.Println(".........wait.................")
	fmt.Println("wait")
	wg.Wait()
	fmt.Println("done")
}