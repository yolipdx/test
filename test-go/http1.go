package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func (w http.ResponseWriter, r * http.Request) {
		fmt.Println("req: ", r)
		io.WriteString(w, "echo for hello")
	})

	
	log.Fatal(http.ListenAndServe(":8080", nil))
}