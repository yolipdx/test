package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, jobs chan int, results chan string) {

	for job := range jobs {
		// fmt.Printf("worker %d work on %d\n", id, job)
		time.Sleep(time.Duration(rand.Int() % 8) * time.Second)
		jobResult := fmt.Sprintf("worker: %02d done: for job: %02d with: %02d", id, job, rand.Int() % 100)
		// fmt.Print(jobResult)
		results <- jobResult
	}
}


func main() {

	jobCount := 100
	
	jobCh := make(chan int, jobCount)
	resultsCh := make(chan string, jobCount)

	for i := 0; i < jobCount; i++ {
		jobCh <- i
	}

	for i := 0; i < 10; i++ {
		go worker(i, jobCh, resultsCh)
	}

	for i := 0; i < jobCount; i++ {
		fmt.Println(<-resultsCh)
	}
}