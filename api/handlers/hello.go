package handlers

import (
	"encoding/json"
	"net/http"
)

type helloResponse struct {
	Hello string `json:"hello"`
}

func HandleHello() func(w http.ResponseWriter, _ *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		res := helloResponse{Hello: "World"}
		b, _ := json.Marshal(res)
		w.Write(b)
	}
}
