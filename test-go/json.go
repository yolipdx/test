package main

import (
	"encoding/json"
	"fmt"
)

type reponse1 struct {
	Page   int
	Fruits []string
}

type reponse2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	r1 := reponse1{
		100,
		[]string{"apple", "banana"},
	}
	r1B, _ := json.Marshal(&r1)
	fmt.Println("r1b: ", string(r1B))

	r2 := reponse2{
		100,
		[]string{"apple", "banana"},
	}
	r2B, _ := json.MarshalIndent(&r2, "", "  ")
	fmt.Println("r2b: ", string(r2B))

}
