package main

import (
	"io"
	"log"
	"net/http"
)

type HomeHandler struct {
	Name string
}

func (hh *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "this is from homehandler")
}

func main0() {
	mux := http.NewServeMux()
	mux.Handle("/home/", &HomeHandler{"home handler"})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("default http handler"))
	})

	srv := http.Server{Addr: ":8080", Handler: mux}
	log.Fatal(srv.ListenAndServe())
}

func main() {

	// x := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("default http handler"))
	// })


	http.Handle("/home/", &HomeHandler{"home handler"})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("default http handler"))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}