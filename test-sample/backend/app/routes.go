package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
}

type S1erver struct {
}


func (s *server) Create() http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/api/v1/test/{message}",  s.handleTest())
	return mux
}



func (s *server) handleTest() http.HandlerFunc {

	type testResponse struct {
		Type string `json:"type"`
		Name string `json:"name"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		data := testResponse{"message", "mike"}

		json.NewEncoder(w).Encode(&data)
	}
}