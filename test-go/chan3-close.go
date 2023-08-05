package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main_1() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			fmt.Printf("%d\n", i)
			ch <- i
		}
		close(ch)
	}()

mainLoop:
	for {
		fmt.Println("\nwaiting")
		select {
		case num, ok := <-ch:
			fmt.Println("Got: ", num, ok)
			if !ok {
				break mainLoop
			}
			// default:
			// 	fmt.Println("got nothing")
		}

		// fmt.Println("next select in main")
		// time.Sleep(time.Second)
	}

	fmt.Println("main done")
}

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			randInt := rand.Int() % 10
			fmt.Println("rand sleep with: ", randInt)
			time.Sleep(time.Second * time.Duration(randInt))

			fmt.Printf("%d\n", i)
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println("\ngot: ", v)
	}

	fmt.Println("main done")
}
