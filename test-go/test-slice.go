package main
import "fmt"

func main() {
	aList := []int{1, 2, 3}

	aList = append(aList, 4)
	fmt.Println(aList)

	// aList = append(aList, 5, "a string")
	// fmt.Println(aList)

	fmt.Println(aList[1:4])
}