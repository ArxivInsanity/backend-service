package main

import (
	_ "github.com/ArxivInsanity/backend-service/docs"
	"github.com/ArxivInsanity/backend-service/src/auth"
	"github.com/ArxivInsanity/backend-service/src/config"
	"github.com/ArxivInsanity/backend-service/src/graph"
	"github.com/ArxivInsanity/backend-service/src/healthcheck"
	"github.com/ArxivInsanity/backend-service/src/paper"
	"github.com/ArxivInsanity/backend-service/src/project"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Arxiv Insanity Backend Service
// @version         1.0
// @description     The backend service for the Arxiv insanity project.
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	log.Debug().Msg("Starting application now")
	err := config.WithDbCon(func(mc *config.MongoCon) {
		router := gin.Default()
		router.Use(config.CORSMiddleware())
		router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		router.GET("/", healthcheck.HealthCheck)
		auth.AddRoutes(router)
		project.AddRoutes(router, mc)
		paper.AddRoutes(router)
		graph.AddRoutes(router)

		router.Run()
	})

	if err != nil {
		log.Error().Msg("Failed to get database connection. Shutting down the application")
		log.Err(err)
		return
	}

}
