package main

import (
	"fmt"
)

type testT struct{}

func (t *testT) value() string {
	return "some thing"
}

func main() {
	var t *testT

	// cha, nil pointer竟然可以call他的method；什么情况~~~~~
	fmt.Println(t.value())
}
