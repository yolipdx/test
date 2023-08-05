package main

import (
	"fmt"
	"time"
)

func test0() {
	fmt.Println("timer0 on")
	t := time.NewTimer(5 * time.Second)
	fmt.Println("timer0 waiting")

	go func() {
		ok := t.Stop()
		fmt.Println("timer stop: ", ok)
	}()

	tValue := <-t.C
	fmt.Println("timer0 expired: ", tValue)
}

func test1() {
	timer := time.NewTimer(time.Second)
	go func() {
		time.Sleep(10 * time.Second)
		stopped := timer.Stop()
		fmt.Println("stop:", stopped)
	}()

	t := <- timer.C
	fmt.Println("timer fired: ", t)

	time.Sleep(2 * time.Second)
}

func test2() {
	ch := make(chan string, 2)

	go func() {
		time.Sleep(time.Second * 5)
		ch <- "gogogo"
	}()

	select {
	case m := <-ch:
		fmt.Println("m got: ", m)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout")
	}
}

func main() {
	// test0()
	test1()
}
