package main

import (
	"fmt"
	"sync"
)

func worker(jobId int) {
	defer func() {
		fmt.Println("done with: ", jobId)
		// wg.Done()
	}()

	fmt.Printf("job %d running\n", jobId)
}

func main() {
	var wg sync.WaitGroup
	jobNum := 3
	
	for i := 0; i< jobNum; i++ {
		id := i
		wg.Add(1)
		
		go func() {
			defer wg.Done()
			worker(id)
		}()
	}

	

	fmt.Println(".........wait.................")
	wg.Wait()
	fmt.Println("done")
}