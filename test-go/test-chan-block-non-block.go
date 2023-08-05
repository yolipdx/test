package main

import "fmt"

func testBlock() {
	// deadlock
	blockChan := make(chan string)

	blockChan <- "message"
	fmt.Printf("msg sent")

	go func() {
		receivedMsg := <- blockChan
		fmt.Println("msg received: ", receivedMsg)
	}();
}

func testBlock1() {
	// deadlock
	blockChan := make(chan string)

	go func() {
		receivedMsg := <- blockChan
		fmt.Println("msg received: ", receivedMsg)
	}();

	blockChan <- "message"
	fmt.Printf("msg sent")
}


func testNonBlock() {
	// deadlock
	blockChan := make(chan string, 100)

	done := make(chan bool)

	blockChan <- "message"
	fmt.Println("msg sent")

	go func() {
		receivedMsg := <- blockChan
		fmt.Println("msg received: ", receivedMsg)


		// receivedMsg1 := <- blockChan
		// receivedMsg2 := <- blockChan
		// receivedMsg3 := <- blockChan
		// receivedMsg4 := <- blockChan
		
		// fmt.Println("msg received: ", receivedMsg1)
		// fmt.Println("msg received: ", receivedMsg2)
		// fmt.Println("msg received: ", receivedMsg3)
		// fmt.Println("msg received: ", receivedMsg4)

		done <- true
	}();

	finished := <- done
	fmt.Println("done: ", finished)
}


func main() {
	testNonBlock()
}
