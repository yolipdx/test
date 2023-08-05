package main
import "fmt"

func main() {
	a := 100;
	
	if b := a * 100; b > 100 {
		fmt.Println("yes");
	} else {
		fmt.Println("no");
	}
}