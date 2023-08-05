package main
import (
	"fmt"
	R "reflect"
)

func main() {
	aList := []int{1, 2, 3, 4, 5}
	bList := []int{1, 2, 3, 4, 5}

	fmt.Println("eq ?", R.DeepEqual(aList, bList))
}