package main

import "fmt"

func main() {
	src := []int{1, 2, 3}

	// var dst []int  ERROR => len of dst should be big enough to hold the items 
	dst := make([]int, len(src))

	fmt.Println("src: ", src)
	n := copy(dst, src)
	fmt.Println("dst: ", n, dst)
	
}