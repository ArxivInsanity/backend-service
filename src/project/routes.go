package project

import (
	"github.com/ArxivInsanity/backend-service/src/database"
	"github.com/gin-gonic/gin"
)

func AddAuthRoutes(router *gin.Engine, mongoCon *database.MongoCon) {
	ph := ProjectHandler{mongoCon.Ctx, mongoCon.Collection}
	router.GET("/api/projects", ph.GetAllProjects)
	router.POST("/api/projects", ph.CreateProject)
	router.PUT("/api/projects/:name", ph.UpdateProject)
	router.DELETE("/api/projects/:name", ph.DeleteProject)

	router.PUT("/api/projects/:name/seedPapers", ph.AddSeedPaper)
	router.DELETE("/api/projects/:name/seedPapers", ph.DeleteSeedPaper)
}
