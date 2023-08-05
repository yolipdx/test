package main
import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	count := 1
	for ;; {
		fmt.Println("count: ", count)
		count += 1;
		if count >= 100 {
			break;
		}
	}
}