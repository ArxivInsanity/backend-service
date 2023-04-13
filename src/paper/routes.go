package paper

import (
	"github.com/gin-gonic/gin"
)

func AddRoutes(router *gin.Engine) {
	router.GET("/api/papers/autocomplete", GetPaperAutocompleteSuggestions)
	router.GET("/api/papers/:id", GetPaperDetails)
}
