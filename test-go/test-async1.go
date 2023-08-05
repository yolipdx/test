package main
import "fmt"

func main() {
	msgChan1 := make(chan string)
	msgChan2 := make(chan string)


	go func(msg string, msgQueue chan<- string) {
		fmt.Println("in: ", msg)
		msgQueue <- msg
	}("m2", msgChan2)


	go func(msg string, msgQueue chan<- string) {
		fmt.Println("in: ", msg)
		msgQueue <- msg
	}("m1", msgChan1)


	// msg1 := <- msgChan1
	// fmt.Println("msg1: ", msg1)

	// msg2 := <- msgChan2
	// fmt.Println("msg2: ", msg2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <- msgChan1:
			fmt.Println("out chan1: ", m1)
		case m2 := <- msgChan2:
			fmt.Println("out chan2: ", m2)
		}
	}

}