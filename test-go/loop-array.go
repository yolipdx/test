package main

import "fmt"


func test() {
	a := [3]int{0, 1, 2}

	for i, v := range(a) {
		a[2] = 10
		if i == 2 {
			fmt.Println(i, v)
		}
	}

	fmt.Println(a)
}

func main() {
	a := []int{1, 3, 5, 7, 9, 11}

	for i, v := range a {
		fmt.Println(i, v)
		a[3] = -100
	}

	fmt.Println()
	test()
}