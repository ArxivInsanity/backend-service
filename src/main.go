package main

import (
	"net/http"

	_ "github.com/ArxivInsanity/backend-service/docs"
	auth "github.com/ArxivInsanity/backend-service/src/auth"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	log.Debug().Msg("Starting application now")
	router := gin.Default()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/healthcheck", healthCheck)

	auth.SetupAuth()
	auth.AddAuthRoutes(router)

	router.Run()
}
