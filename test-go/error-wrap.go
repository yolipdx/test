package main

import (
	"errors"
	"fmt"
)

func main() {
	e1 := errors.New("this is a test err")
	e2 := fmt.Errorf("%w: something is wrong", e1)

	fmt.Println(errors.Is(e2, e1))
	fmt.Printf("e1: %v\n", e1)
	fmt.Printf("e2: %v\n", e2)
}