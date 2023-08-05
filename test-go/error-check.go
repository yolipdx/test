package main

import (
	"errors"
	"fmt"
)

type CustomerError struct {
	info string
}

func (c CustomerError) Error() string {
	return c.info
}

func main() {
	// e := errors.New("this is a error")
	e := CustomerError{"test customer error"}
	e2 := fmt.Errorf("%w wrapped in e2", e)

	switch {
	case errors.As(e2, &CustomerError{}):
		fmt.Println("customer error")
	default:
		break
	}

	errors.
}