package main

import (
	_ "github.com/ArxivInsanity/backend-service/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// Health Check godoc
// @Summary Endpoint for Health Check
// @Schemes
// @Description Endpoint for performing health check on the application
// @Tags Health Check
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Router /healthcheck [get]
func healthCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Up")
}

// @title           Arxiv Insanity Backend Service
// @version         1.0
// @description     The backend service for the Arxiv insanity project.
func main() {
	router := gin.Default()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/healthcheck", healthCheck)
	router.Run()
}
