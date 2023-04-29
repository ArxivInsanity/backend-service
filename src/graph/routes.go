package graph

import (
	"github.com/gin-gonic/gin"
)

func AddRoutes(router *gin.Engine) {
	router.GET("/api/graph/:id", GetGraph)
}
