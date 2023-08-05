package main

import (
	// "errors"
	"log"
	"strings"
)

type MultiError struct {
	errorList []string
}

func (me *MultiError) Add(err error) {
	me.errorList = append(me.errorList, err.Error())
}

func (me MultiError) Error() string {
	return strings.Join(me.errorList, ":")
}

type Customer struct {
	name string
	age  int
}

func (c Customer) Validate0() *MultiError {
	var m *MultiError
	// if c.age < 0 {
	// 	m = &MultiError{}
	// 	m.Add(errors.New("age is negative"))
	// }

	// if c.name == "" {
	// 	if m == nil {
	// 		m = &MultiError{}
	// 	}
	// 	m.Add(errors.New("name is nil"))
	// }

	return m
}


func (c Customer) Validate1() error {
	var m *MultiError
	// if c.age < 0 {
	// 	m = &MultiError{}
	// 	m.Add(errors.New("age is negative"))
	// }

	// if c.name == "" {
	// 	if m == nil {
	// 		m = &MultiError{}
	// 	}
	// 	m.Add(errors.New("name is nil"))
	// }

	return m
}

func main() {
	c := Customer{"jk", 100}
	if e := c.Validate0(); e != nil {
		log.Fatalf("v0: customer is invalid: %v", e)
	}

	if e := c.Validate1(); e != nil {
		log.Fatalf("v1: customer is invalid: %v", e)
	}
}
