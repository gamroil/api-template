package handlers

import (
	"api-template/api/metrics"
	"encoding/json"
	"net/http"
)

type fooResponse struct {
	Foo string `json:"foo"`
}

func HandleFoo(metrics *metrics.Client) func(w http.ResponseWriter, _ *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		metrics.IncrementFooCount()
		res := fooResponse{Foo: "Bar"}
		b, _ := json.Marshal(res)
		_, _ = w.Write(b)
	}
}
