package main

import "fmt"

func Map[T1 any, T2 any](mapper func(T1) T2, values []T1) []T2 {
	valuesLen := len(values)
	mappedValues := make([]T2, valuesLen)
	for _, v := range values {
		mappedValues = append(mappedValues, mapper(v))
	}

	return mappedValues
}

func main() {
	nums := []int{1, 3, 5, 7, 9}

	s := Map(func(v int) string {
		fmt.Println("v: ", v)
		return fmt.Sprintf("AS %d", v)
	}, nums)

	fmt.Println(s)
}