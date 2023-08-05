package main

// better solutions at 
// https://gist.github.com/kaipakartik/8120855
import (
	"fmt"
	"golang.org/x/tour/tree"
	// "time"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value

	if t.Left != nil {
		Walk(t.Left, ch)
	}

	if t.Right != nil {
		Walk(t.Right, ch)
	}

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 20)
	ch2 := make(chan int, 20)

	go Walk(tree.New(1), ch1)
	go Walk(tree.New(1), ch2)

	vList1 := []int{}
	vList2 := []int{}

	for i := 0; i < 10; i++ {
		vList1 = append(vList1, <-ch1)
	}

	for i:=0; i< 10; i++ {
		vList2 = append(vList2, <-ch2)
	}
	fmt.Println("vList1: ", vList1)
	fmt.Println("vList2: ", vList2)

	if len(vList1) != len(vList2) {
		return false
	}
	
	for i:=0; i< len(vList1); i++ {
		if vList1[i] != vList2[i] {
			return false;
		}
	}

	return true;
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(1)

	fmt.Println(Same(t1, t2))

}
