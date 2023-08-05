package main
import "fmt"

func main() {
	m := make(map[string]int)	

	m["hi"] = 100

	v, inMap := m["hi0"]

	fmt.Println(m["hi"])
	fmt.Println("v: ", v)
	fmt.Println("inMap: ", inMap)
}