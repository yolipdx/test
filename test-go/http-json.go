package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Profile struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Blogs    []Blog `json:"blogs,omitempty"`
}

type Blog struct {
	BlogName string `json:"name"`
	URL      string `json:"url"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		profile := Profile{Email: "abhirockzz@gmail.com", Username: "abhirockzz", Blogs: []Blog{
			{BlogName: "devto", URL: "https://dev.to/abhirockzz/"},
			{BlogName: "medium", URL: "https://medium.com/@abhishek1987/"},
		}}

		w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")
		encoder := json.NewEncoder(w)
		encoder.Encode(&profile)

		fmt.Println("reponsed with: ", profile)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
