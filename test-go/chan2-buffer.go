package main

import "fmt"

func main_1() {
	ch := make(chan int, 10)

	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4

	fmt.Println(" main done: not blocked")


	chBlocked := make(chan int)
	chBlocked <- 1
	chBlocked <- 2
	chBlocked <- 3
	chBlocked <- 4

	fmt.Println("!! main done: will not reach here !!")
}

func main() {
	ch := make(chan int, 2)

	ch <- 1
	fmt.Println("not blocked yet")
	ch <- 2
	fmt.Println("not blocked yet")
	ch <- 3
	fmt.Println("blocked <- full")
	ch <- 4
	fmt.Println("blocked <- full")

	fmt.Println(" main done: blocked now")
}