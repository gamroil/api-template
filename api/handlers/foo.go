package handlers

import (
	"api-template/api/observability"
	"net/http"

	"github.com/gin-gonic/gin"
)

type fooResponse struct {
	Foo string `json:"foo"`
}

func HandleFoo(metrics *observability.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		metrics.IncrementFooCount()
		res := fooResponse{Foo: "Bar"}
		c.JSON(http.StatusOK, res)
	}
}
