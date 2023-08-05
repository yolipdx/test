package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func c1(ctx context.Context, id int) {


	for i := 0; i < 2; i++ {
		go c2(ctx, i * id)
	}

	for {
		select {
		case <- ctx.Done():
			fmt.Printf("c1 %d end\n", id)
			return
		default:
			// fmt.Printf("c1 %d tick\n", id)
			time.Sleep(time.Duration(rand.Int() % 10) * time.Second)
		}
	}
}

func c2(ctx context.Context, id int) {


	for i := 0; i < 2; i++ {
		go c3(ctx, i * id)
	}

	for {
		select {
		case <- ctx.Done():
			fmt.Printf("c2 %d end\n", id)
			return
		default:
			// fmt.Printf("c2 %d tick\n", id)
			time.Sleep(time.Duration(rand.Int() % 10) * time.Second)
		}
	}
}


func c3(ctx context.Context, id int) {

	for {
		select {
		case <- ctx.Done():
			fmt.Printf("c3 %d end\n", id)
			return
		default:
			// fmt.Printf("c3 %d tick\n", id)
			time.Sleep(time.Duration(rand.Int() % 10) * time.Second)
		}
	}
}


func main() {

	// opttion 1
	// ctx, cancel := context.WithCancel(context.Background())


	// opttion 2
	// deadline := time.Now().Add(5 * time.Second)
	// ctx, cancel := context.WithDeadline(context.Background(), deadline)
	// defer cancel()


	// opttion 3
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	ctx.
	defer cancel()

	for i := 0; i < 4; i++ {
		go c1(ctx, i)
	}

	
	time.Sleep(time.Second * 5)
	fmt.Printf("num of threads: %d\n", runtime.NumGoroutine())
	fmt.Println("cancel incoming....")
	// cancel()
	time.Sleep(time.Second * 10)
	fmt.Printf("num of threads: %d\n", runtime.NumGoroutine())
	fmt.Println("main done")
}