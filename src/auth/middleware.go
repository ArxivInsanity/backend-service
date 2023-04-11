package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func AuthenticationRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Debug().Msg("Validating user creds")
		status := IsValidLogin(c)
		if !status {
			c.AbortWithStatusJSON(http.StatusForbidden, "User not logged in")
			log.Error().Msg("User is not logged in")
			return
		}
		log.Debug().Msg("User is logged in. Okay to proceed")
		c.Set(USER, GetUserEmail(c))
		c.Next()
	}
}
