package main

import "fmt"

func main() {
	s := "hêllo中国"

	for i, v := range s {
		fmt.Printf("%d %c %c\n", i, v, rune(s[i]))
	}
}