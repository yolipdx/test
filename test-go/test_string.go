package main
import (
	f "fmt"
	s "strings"
)


var p = f.Println

func main() {
	p("contains: ", s.Contains("test", "es"))
}