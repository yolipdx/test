package main
import "fmt"

func main() {
	switch i := 100; i / 1000 {
	case 100:
		fmt.Println("100")

	default:
		fmt.Println("default")
		break
	}

	testV := -100
	switch {
	case testV < 100:
		fmt.Println("< 100")
	case testV == 100:
		fmt.Println("== 100")
	default:
		fmt.Println("<100")
	}
}