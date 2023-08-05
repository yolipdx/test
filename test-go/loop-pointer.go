package main

import "fmt"

type people struct {
	name string
	age  int
}

func main() {
	peopleList := [4]people{
		{"p0", 10},
		{"p1", 30},
		{"p2", 20},
		{"p3", 60},
	}

	another := []people{}
	for _, p := range peopleList {
		fmt.Printf("%p:%+v\n", &p, p)

		another = append(another, p)
	}

	fmt.Println(another)


	nums := []int{1, 3, 4, 5}
	for _, v := range nums {
		fmt.Printf("%p: %d\n", &v, v)
	}
}
