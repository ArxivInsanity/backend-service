package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func healthCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Up")
}

func main() {
	router := gin.Default()
	router.GET("/healthcheck", healthCheck)
	router.Run()
}
