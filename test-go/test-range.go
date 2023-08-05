package main
import "fmt"

var p = fmt.Println

func main() {
	aArray := [10]int{3, 1, 123, -110}
	aList :=  []int{6, 1, 3, -11}
	aMap := map[string]int {
		"hi": 100,
		"age": -100,
	}

	for i, v := range aArray {
		p(i, v)
	}
	p("")

	for i, v := range aList {
		p(i, v)
	}
	p("")

	for i, v := range aMap {
		p(i, v)
	}
	p("")

	for i, v := range "this is a tring" {
		p(i, v)
	}
}