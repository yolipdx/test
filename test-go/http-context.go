package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func handleHello(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	io.WriteString(w, "hello")
	
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("time out")
	case <- ctx.Done():
		err := ctx.Err()
		fmt.Println("http error: ", err)
		// status := http.StatusInternalServerError
		// http.Error(w, err.Error(), status)
		io.WriteString(w, "hello")
	}

}

func main() {
	http.HandleFunc("/hello", handleHello)

	http.ListenAndServe(":8090", nil)
}