package main

import (
	"fmt"
	"os"
)

func main() {
	ch := make(chan string, 6)
	ch <- "./test-files/1.txt"
	ch <- "./test-files/2.txt"
	ch <- "./test-files/34.txt"
	ch <- "./test-files/4.txt"
	close(ch)

	for path := range ch {
		func() {
			f, err := os.Open(path)
			if err != nil {
				fmt.Println("err: ", err)
				return
			}

			defer f.Close()
		}()
	}
}