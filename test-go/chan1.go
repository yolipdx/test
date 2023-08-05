package main

import (
	"fmt"
	"time"
)

func doSum(nums []int, messageChan chan int, done chan bool) {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println(nums, sum)
	messageChan <- sum

	done <- true
}

func main() {
	nums := []int{3, 1, -123, 123, 3221}

	// fmt.Println((nums))

	message := make(chan int)
	done := make(chan bool)

	go doSum(nums[:len(nums)/2], message, done)
	go doSum(nums[len(nums)/2:], message, done)

	// v1 := <-message
	// v2 := <-message

	// fmt.Println("v1: ", v1)
	// fmt.Println("v2: ", v2)

	for  {
		isDone := false
		select {
		case v := <-message:
			fmt.Println("got one: ", v)
		case isDone = <-done:
			fmt.Println("got done: ", isDone)
			if isDone {
				break
			}
		default:
			fmt.Println("nothing")
		}

		time.Sleep(time.Second)
		if isDone {
			break
		}
	}

	fmt.Println("main done")
}
