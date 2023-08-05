package main
import "fmt"



func add(nums ...int) int {
	result := 0

	for _, v := range nums {
		result += v
	}

	fmt.Println("result: ", result)
	return result
}

func main() {
	// a := []int{1, 2, 3, 4}
	a := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("%d\n", add(a...))

}