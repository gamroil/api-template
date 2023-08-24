package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type helloResponse struct {
	Hello string `json:"hello"`
}

func HandleHello() func(c *gin.Context) {
	return func(c *gin.Context) {
		res := helloResponse{Hello: "World"}
		c.JSON(http.StatusOK, res)
	}
}
