package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health Check godoc
// @Summary Endpoint for Health Check
// @Schemes
// @Description Endpoint for performing health check on the application
// @Tags Health Check
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Router / [get]
func HealthCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Up")
}
