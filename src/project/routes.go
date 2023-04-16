package project

import (
	"github.com/ArxivInsanity/backend-service/src/config"
	"github.com/gin-gonic/gin"
)

func AddRoutes(router *gin.Engine, mongoCon *config.MongoCon) {
	ph := ProjectHandler{mongoCon.Ctx, mongoCon.Collection}
	router.GET("/api/projects", ph.GetAllProjects)
	router.POST("/api/projects", ph.CreateProject)
	router.PUT("/api/projects/:name", ph.UpdateProject)
	router.DELETE("/api/projects/:name", ph.DeleteProject)

	router.GET("/api/projects/:name/seedPapers", ph.GetSeedPapers)
	router.PUT("/api/projects/:name/seedPapers", ph.AddSeedPaper)
	router.DELETE("/api/projects/:name/seedPapers", ph.DeleteSeedPaper)
}
