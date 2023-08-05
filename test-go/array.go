package main
import "fmt"

func main() {
	aList := [20]int{1, 2, 3, 4}
	bList := []int{}
	cList := [3][10]int{
		{1, 2, 3},
	}

	// aList = append(aList, -10)

	fmt.Println(aList)
	fmt.Println(bList)

	// bList[10] = 100
	// bList.append(100)

	fmt.Printf("len of aList: %d\n", len(aList))
	fmt.Printf("len of cList: %d\n", len(cList))
	fmt.Printf("cList: ", cList)
}