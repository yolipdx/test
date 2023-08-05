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

	another := []*people{}
	for id, p := range peopleList {
		p.name = "ho" + fmt.Sprintf("name with id: %d", id)
		fmt.Printf("%p:%+v\n", &p, p)

		another = append(another, &p)
	}

	fmt.Println(peopleList)
	fmt.Println(another)
}
